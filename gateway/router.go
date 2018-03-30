package gateway

import (
	stdctx "context"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/ansel1/merry"

	"github.com/percolate/shisa/context"
	"github.com/percolate/shisa/httpx"
	"github.com/percolate/shisa/metrics"
	"github.com/percolate/shisa/service"
)

const (
	// RequestIdGenerationMetricKey is the `ResponseSnapshot` metric for generating the request id
	RequestIdGenerationMetricKey = "request-id-generation"
	// ParseQueryParametersMetricKey is the `ResponseSnapshot` metric for parsing the request query parameters
	ParseQueryParametersMetricKey = "parse-query"
	// RunGatewayHandlersMetricKey is the `ResponseSnapshot` metric for running the Gateway level handlers
	RunGatewayHandlersMetricKey = "handlers"
	// FindEndpointMetricKey is the `ResponseSnapshot` metric for resolving the request's endpoint
	FindEndpointMetricKey = "find-endpoint"
	// ValidateQueryParametersMetricKey is the `ResponseSnapshot` metric for validating the request query parameters
	ValidateQueryParametersMetricKey = "validate-query"
	// RunEndpointPipelineMetricKey is the `ResponseSnapshot` metric for running the endpoint's pipeline
	RunEndpointPipelineMetricKey = "pipeline"
	// SerializeResponseMetricKey is the `ResponseSnapshot` metric for serializing the response
	SerializeResponseMetricKey = "serialization"
)

var (
	timingPool = sync.Pool{
		New: func() interface{} {
			return metrics.NewTiming()
		},
	}
)

func getTiming() *metrics.Timing {
	timing := timingPool.Get().(*metrics.Timing)
	timing.ResetAll()

	return timing
}

func putTiming(timing *metrics.Timing) {
	timingPool.Put(timing)
}

func (g *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ri := httpx.NewInterceptor(w)

	ctx := context.Get(r.Context())
	defer context.Put(ctx)

	request := httpx.GetRequest(r)
	defer httpx.PutRequest(request)

	timing := getTiming()
	defer putTiming(timing)

	timing.Start(RequestIdGenerationMetricKey)
	requestID, idErr := g.generateRequestID(ctx, request)
	timing.Stop(RequestIdGenerationMetricKey)

	ctx = ctx.WithRequestID(requestID)
	ri.Header().Set(g.RequestIDHeaderName, requestID)

	timing.Start(ParseQueryParametersMetricKey)
	parseOK := request.ParseQueryParameters()
	timing.Stop(ParseQueryParametersMetricKey)

	var cancel stdctx.CancelFunc
	ctx, cancel = ctx.WithCancel()
	defer cancel()

	if cn, ok := w.(http.CloseNotifier); ok {
		go func() {
			select {
			case <-cn.CloseNotify():
				cancel()
			case <-ctx.Done():
			}
		}()
	}

	var (
		path       string
		endpoint   *endpoint
		pipeline   *service.Pipeline
		err        merry.Error
		response   httpx.Response
		tsr        bool
		responseCh chan httpx.Response = make(chan httpx.Response, 1)
	)

	subCtx := ctx
	if g.HandlersTimeout != 0 {
		var subCancel stdctx.CancelFunc
		subCtx, subCancel = subCtx.WithTimeout(g.HandlersTimeout - ri.Elapsed())
		defer subCancel()
	}

	timing.Start(RunGatewayHandlersMetricKey)
	for _, handler := range g.Handlers {
		go func() {
			response, exception := handler.InvokeSafely(subCtx, request)
			if exception != nil {
				err = exception.Prepend("gateway: route: run gateway handler")
				response = g.handleError(subCtx, request, err)
			}

			responseCh <- response
		}()
		select {
		case <-subCtx.Done():
			timing.Stop(RunGatewayHandlersMetricKey)
			cancel()
			err = merry.Prepend(subCtx.Err(), "gateway: route: request aborted")
			if merry.Is(subCtx.Err(), stdctx.DeadlineExceeded) {
				err = err.WithHTTPCode(http.StatusGatewayTimeout)
			}
			response = g.handleError(subCtx, request, err)
			goto finish
		case response = <-responseCh:
			if response != nil {
				timing.Stop(RunGatewayHandlersMetricKey)
				goto finish
			}
		}
	}
	timing.Stop(RunGatewayHandlersMetricKey)

	timing.Start(FindEndpointMetricKey)
	path = request.URL.EscapedPath()
	endpoint, request.PathParams, tsr, err = g.tree.getValue(path)
	timing.Stop(FindEndpointMetricKey)

	if err != nil {
		err = err.Prepend("gateway: route")
		response = g.handleError(ctx, request, err)
		goto finish
	}

	if endpoint == nil {
		response, err = g.handleNotFound(ctx, request)
		goto finish
	}

	switch request.Method {
	case http.MethodHead:
		pipeline = endpoint.Head
	case http.MethodGet:
		pipeline = endpoint.Get
	case http.MethodPut:
		pipeline = endpoint.Put
	case http.MethodPost:
		pipeline = endpoint.Post
	case http.MethodPatch:
		pipeline = endpoint.Patch
	case http.MethodDelete:
		pipeline = endpoint.Delete
	case http.MethodConnect:
		pipeline = endpoint.Connect
	case http.MethodOptions:
		pipeline = endpoint.Options
	case http.MethodTrace:
		pipeline = endpoint.Trace
	}

	if pipeline == nil {
		if tsr {
			response, err = g.handleNotFound(ctx, request)
		} else {
			response, err = endpoint.handleNotAllowed(ctx, request)
		}
		goto finish
	}

	if tsr {
		if path != "/" && pipeline.Policy.AllowTrailingSlashRedirects {
			response, err = endpoint.handleRedirect(ctx, request)
		} else {
			response, err = g.handleNotFound(ctx, request)
		}
		goto finish
	}

	if !parseOK && !pipeline.Policy.AllowMalformedQueryParameters {
		response, err = endpoint.handleBadQuery(ctx, request)
		goto finish
	}

	timing.Start(ValidateQueryParametersMetricKey)
	if malformed, unknown, exception := request.ValidateQueryParameters(pipeline.QueryFields); exception != nil {
		timing.Stop(ValidateQueryParametersMetricKey)
		response, exception = endpoint.handleError(ctx, request, exception)
		if exception != nil {
			g.invokeErrorHookSafely(ctx, request, exception)
		}
		goto finish
	} else if malformed && !pipeline.Policy.AllowMalformedQueryParameters {
		timing.Stop(ValidateQueryParametersMetricKey)
		response, err = endpoint.handleBadQuery(ctx, request)
		goto finish
	} else if unknown && !pipeline.Policy.AllowUnknownQueryParameters {
		timing.Stop(ValidateQueryParametersMetricKey)
		response, err = endpoint.handleBadQuery(ctx, request)
		goto finish
	}
	timing.Stop(ValidateQueryParametersMetricKey)

	if !pipeline.Policy.PreserveEscapedPathParameters {
		for i := range request.PathParams {
			if esc, r := url.PathUnescape(request.PathParams[i].Value); r == nil {
				request.PathParams[i].Value = esc
			}
		}
	}

	if pipeline.Policy.TimeBudget != 0 {
		var cancel stdctx.CancelFunc
		ctx, cancel = ctx.WithTimeout(pipeline.Policy.TimeBudget - ri.Elapsed())
		defer cancel()
	}

	timing.Start(RunEndpointPipelineMetricKey)
	select {
	case <-ctx.Done():
		timing.Stop(RunEndpointPipelineMetricKey)
		err = merry.Prepend(ctx.Err(), "gateway: route: request aborted")
		if merry.Is(ctx.Err(), stdctx.DeadlineExceeded) {
			err = err.WithHTTPCode(http.StatusGatewayTimeout)
		}
		response = g.handleEndpointError(endpoint, ctx, request, err)
		goto finish
	default:
	}

endpointHandlers:
	for _, handler := range pipeline.Handlers {
		go func() {
			response, exception := handler.InvokeSafely(ctx, request)
			if exception != nil {
				err = exception.Prepend("gateway: route: run endpoint handler")
				response = g.handleEndpointError(endpoint, ctx, request, err)
			}

			responseCh <- response
		}()
		select {
		case <-ctx.Done():
			timing.Stop(RunEndpointPipelineMetricKey)
			err = merry.Prepend(ctx.Err(), "gateway: route: request aborted")
			if merry.Is(ctx.Err(), stdctx.DeadlineExceeded) {
				err = err.WithHTTPCode(http.StatusGatewayTimeout)
			}
			response = g.handleEndpointError(endpoint, ctx, request, err)
			goto finish
		case response = <-responseCh:
			if response != nil {
				break endpointHandlers
			}
		}
	}
	timing.Stop(RunEndpointPipelineMetricKey)

	if response == nil {
		err = merry.New("gateway: route: no response from pipeline")
		response = g.handleEndpointError(endpoint, ctx, request, err)
	}

finish:
	timing.Start(SerializeResponseMetricKey)
	var (
		writeErr merry.Error
		snapshot httpx.ResponseSnapshot
	)
	if merry.Is(ctx.Err(), stdctx.Canceled) {
		writeErr = merry.New("gateway: route: user agent disconnect")
		snapshot = ri.Snapshot()
	} else {
		writeErr = ri.WriteResponse(response)
		writeErr = merry.Prepend(writeErr, "gateway: route: serialize response")
		snapshot = ri.Flush()
	}
	timing.Stop(SerializeResponseMetricKey)

	if g.CompletionHook != nil {
		timing.Do(func(name string, timer *metrics.Timer) {
			snapshot.Metrics[name] = timer.Interval()
		})

		g.invokeCompletionHookSafely(ctx, request, snapshot)
	}

	if idErr != nil {
		g.invokeErrorHookSafely(ctx, request, idErr)
	}

	if err != nil {
		g.invokeErrorHookSafely(ctx, request, err)
	}

	if writeErr != nil {
		g.invokeErrorHookSafely(ctx, request, writeErr)
	}

	respErr := response.Err()
	if respErr != nil && respErr != err {
		respErr1 := merry.Prepend(respErr, "gateway: route: handler failed")
		g.invokeErrorHookSafely(ctx, request, respErr1)
	}
}

func (g *Gateway) generateRequestID(ctx context.Context, request *httpx.Request) (string, merry.Error) {
	if g.RequestIDGenerator == nil {
		return request.ID(), nil
	}

	requestID, err, exception := g.RequestIDGenerator.InvokeSafely(ctx, request)
	if exception != nil {
		err = exception.Prepend("gateway: route: generate request id")
		requestID = request.ID()
	} else if err != nil {
		err = err.Prepend("gateway: route: generate request id")
		requestID = request.ID()
	} else if requestID == "" {
		err = merry.New("gateway: route: generate request id: empty value")
		requestID = request.ID()
	}

	return requestID, err
}

func (g *Gateway) handleNotFound(ctx context.Context, request *httpx.Request) (httpx.Response, merry.Error) {
	if g.NotFoundHandler == nil {
		return httpx.NewEmpty(http.StatusNotFound), nil
	}

	response, exception := g.NotFoundHandler.InvokeSafely(ctx, request)
	if exception != nil {
		err := exception.Prepend("gateway: route: run NotFoundHandler")
		return httpx.NewEmpty(http.StatusNotFound), err
	}

	return response, nil
}

func (g *Gateway) handleError(ctx context.Context, request *httpx.Request, err merry.Error) httpx.Response {
	if g.InternalServerErrorHandler == nil {
		return httpx.NewEmptyError(merry.HTTPCode(err), err)
	}

	response, exception := g.InternalServerErrorHandler.InvokeSafely(ctx, request, err)
	if exception != nil {
		response = httpx.NewEmptyError(merry.HTTPCode(err), err)
		exception = exception.Prepend("gateway: route: run InternalServerErrorHandler")
		g.invokeErrorHookSafely(ctx, request, exception)
	}

	return response
}

func (g *Gateway) handleEndpointError(endpoint *endpoint, ctx context.Context, request *httpx.Request, err merry.Error) httpx.Response {
	response, exception := endpoint.handleError(ctx, request, err)
	if exception != nil {
		g.invokeErrorHookSafely(ctx, request, exception)
	}

	return response
}

func (g *Gateway) invokeErrorHookSafely(ctx context.Context, request *httpx.Request, err merry.Error) {
	if g.ErrorHook == nil {
		log.Println(ctx.RequestID(), err.Error())
	}

	if exception := g.ErrorHook.InvokeSafely(ctx, request, err); exception != nil {
		log.Println(ctx.RequestID(), err.Error())
		exception = exception.Prepend("gateway: route: run ErrorHook")
		log.Println(ctx.RequestID(), exception.Error())
	}
}

func (g *Gateway) invokeCompletionHookSafely(ctx context.Context, request *httpx.Request, snapshot httpx.ResponseSnapshot) {
	if exception := g.CompletionHook.InvokeSafely(ctx, request, snapshot); exception != nil {
		exception = exception.Prepend("gateway: route: run CompletionHook")
		g.invokeErrorHookSafely(ctx, request, exception)
	}
}
