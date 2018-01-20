package service

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/percolate/shisa/context"
)

// Handler is a block of logic to apply to a request.
// Returning a non-nil response indicates further request
// processing should be stopped.
type Handler func(context.Context, *Request) Response

// Pipeline is a chain of handlers to be invoked in order on a
// request.  The first non-nil response will be returned to the
// user agent.  If no response is produced an Internal Service
// Error handler will be invoked.
type Pipeline struct {
	Policy      Policy    // customizes automated behavior
	Handlers    []Handler // the pipline steps, minimum one
	QueryFields []Field   // optional query parameter validation
}

// Endpoint is collection of pipelines for a route (URL path),
// one for each HTTP method.  Only supported methods should have
// pipelines, but at least one pipleline is requried.
type Endpoint struct {
	Route   string    // the absolute URL path for this endpoint
	Head    *Pipeline // HEAD method pipeline
	Get     *Pipeline // GET method pipeline
	Put     *Pipeline // PUT method pipeline
	Post    *Pipeline // POST method pipeline
	Patch   *Pipeline // PATCH method pipeline
	Delete  *Pipeline // DELETE method pipeline
	Connect *Pipeline // CONNECT method pipeline
	Options *Pipeline // OPTIONS method pipeline
	Trace   *Pipeline // TRACE method pipeline
}

// GetEndpoint returns an Endpoint configured for the given route
// with the GET pipeline using the given handlers.
func GetEndpoint(route string, handlers ...Handler) Endpoint {
	return Endpoint{
		Route: route,
		Get: &Pipeline{
			Handlers: handlers,
		},
	}
}

// GetEndpointWithPolicy returns an Endpoint configured for the
// given route, with the given policy and the GET pipeline using
// the given handlers.
func GetEndpointWithPolicy(route string, policy Policy, handlers ...Handler) Endpoint {
	return Endpoint{
		Route: route,
		Get: &Pipeline{
			Policy:   policy,
			Handlers: handlers,
		},
	}
}

// PutEndpoint returns an Endpoint configured for the given route
// with the PUT pipeline using the given handlers.
func PutEndpoint(route string, handlers ...Handler) Endpoint {
	return Endpoint{
		Route: route,
		Put: &Pipeline{
			Handlers: handlers,
		},
	}
}

// PutEndpointWithPolicy returns an Endpoint configured for the
// given route, with the given policy and the PUT pipeline using
// the given handlers.
func PutEndpointWithPolicy(route string, policy Policy, handlers ...Handler) Endpoint {
	return Endpoint{
		Route: route,
		Put: &Pipeline{
			Policy:   policy,
			Handlers: handlers,
		},
	}
}

// PostEndpoint returns an Endpoint configured for the given route
// with the POST pipeline using the given handlers.
func PostEndpoint(route string, handlers ...Handler) Endpoint {
	return Endpoint{
		Route: route,
		Post: &Pipeline{
			Handlers: handlers,
		},
	}
}

// PostEndpointWithPolicy returns an Endpoint configured for the
// given route, with the given policy and the POST pipeline
// using the given handlers.
func PostEndpointWithPolicy(route string, policy Policy, handlers ...Handler) Endpoint {
	return Endpoint{
		Route: route,
		Post: &Pipeline{
			Policy:   policy,
			Handlers: handlers,
		},
	}
}

// PatchEndpoint returns an Endpoint configured for the given
// route with the PATCH pipeline using the given handlers.
func PatchEndpoint(route string, handlers ...Handler) Endpoint {
	return Endpoint{
		Route: route,
		Patch: &Pipeline{
			Handlers: handlers,
		},
	}
}

// PatchEndpointWithPolicy returns an Endpoint configured for the
// given route, with the given policy and the PATCH pipeline
// using the given handlers.
func PatchEndpointWithPolicy(route string, policy Policy, handlers ...Handler) Endpoint {
	return Endpoint{
		Route: route,
		Patch: &Pipeline{
			Policy:   policy,
			Handlers: handlers,
		},
	}
}

// DeleteEndpoint returns an Endpoint configured for the given
// route with the DELETE pipeline using the given handlers.
func DeleteEndpoint(route string, handlers ...Handler) Endpoint {
	return Endpoint{
		Route: route,
		Delete: &Pipeline{
			Handlers: handlers,
		},
	}
}

// DeleteEndpointWithPolicy returns an Endpoint configured for
// the given route, with the given policy and the DELETE
// pipeline using the given handlers.
func DeleteEndpointWithPolicy(route string, policy Policy, handlers ...Handler) Endpoint {
	return Endpoint{
		Route: route,
		Delete: &Pipeline{
			Policy:   policy,
			Handlers: handlers,
		},
	}
}

// String implements `expvar.Var.String`
func (e Endpoint) String() string {
	var buf bytes.Buffer
	var written bool
	buf.WriteByte('{')
	if e.Head != nil {
		written = true
		buf.WriteString("\"HEAD\":")
		e.Head.jsonify(&buf)
	}
	if e.Get != nil {
		if written {
			buf.WriteByte(',')
		} else {
			written = true
		}
		buf.WriteString("\"GET\":")
		e.Get.jsonify(&buf)
	}
	if e.Put != nil {
		if written {
			buf.WriteByte(',')
		} else {
			written = true
		}
		buf.WriteString("\"PUT\":")
		e.Put.jsonify(&buf)
	}
	if e.Post != nil {
		if written {
			buf.WriteByte(',')
		} else {
			written = true
		}
		buf.WriteString("\"POST\":")
		e.Post.jsonify(&buf)
	}
	if e.Patch != nil {
		if written {
			buf.WriteByte(',')
		} else {
			written = true
		}
		buf.WriteString("\"PATCH\":")
		e.Patch.jsonify(&buf)
	}
	if e.Delete != nil {
		if written {
			buf.WriteByte(',')
		} else {
			written = true
		}
		buf.WriteString("\"DELETE\":")
		e.Delete.jsonify(&buf)
	}
	if e.Connect != nil {
		if written {
			buf.WriteByte(',')
		} else {
			written = true
		}
		buf.WriteString("\"CONNECT\":")
		e.Connect.jsonify(&buf)
	}
	if e.Options != nil {
		if written {
			buf.WriteByte(',')
		} else {
			written = true
		}
		buf.WriteString("\"OPTIONS\":")
		e.Options.jsonify(&buf)
	}
	if e.Trace != nil {
		if written {
			buf.WriteByte(',')
		}
		buf.WriteString("\"TRACE\":")
		e.Trace.jsonify(&buf)
	}
	buf.WriteByte('}')

	return buf.String()
}

func (p Pipeline) jsonify(buf *bytes.Buffer) {
	enc := json.NewEncoder(buf)
	buf.WriteString("{\"Policy\":")
	enc.Encode(p.Policy)
	buf.WriteString(",\"Handlers\":")
	buf.WriteString(strconv.Itoa(len(p.Handlers)))
	if len(p.QueryFields) != 0 {
		buf.WriteString(",\"QueryFields\":")
		enc.Encode(p.QueryFields)
	}
	buf.WriteByte('}')
}
