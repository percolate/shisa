package main

import (
	"net/http"
	"time"

	"github.com/ansel1/merry"

	"github.com/percolate/shisa/context"
	"github.com/percolate/shisa/env"
	"github.com/percolate/shisa/middleware"
	"github.com/percolate/shisa/service"
)

type GoodbyeService struct {
	service.ServiceAdapter
	env       env.Provider
	endpoints []service.Endpoint
}

func NewGoodbyeService(environment env.Provider) *GoodbyeService {
	policy := service.Policy{
		TimeBudget:                  time.Millisecond * 5,
		AllowTrailingSlashRedirects: true,
	}

	svc := &GoodbyeService{
		env: environment,
	}

	proxy := middleware.ReverseProxy{
		Router:    svc.router,
		Responder: svc.responder,
	}
	farewell := service.GetEndpointWithPolicy("/api/farewell", policy, proxy.Service)
	farewell.Get.QueryFields = []service.Field{{Name: "name", Multiplicity: 1}}

	svc.endpoints = []service.Endpoint{farewell}

	return svc
}

func (s *GoodbyeService) Name() string {
	return "goodbye"
}

func (s *GoodbyeService) Endpoints() []service.Endpoint {
	return s.endpoints
}

func (s *GoodbyeService) Healthcheck() merry.Error {
	addr, envErr := s.env.Get(goodbyeServiceAddrEnv)
	if envErr != nil {
		return envErr.WithUserMessage("address environment variable not found")
	}

	response, err := http.Get("http://" + addr + "/healthcheck")
	if err != nil {
		return merry.Wrap(err).WithUserMessage("unable to complete request")
	}
	if response.StatusCode != http.StatusOK {
		return merry.New("not ready").WithUserMessage("not ready")
	}

	return nil
}

func (s *GoodbyeService) router(ctx context.Context, request *service.Request) (*service.Request, merry.Error) {
	addr, envErr := s.env.Get(goodbyeServiceAddrEnv)
	if envErr != nil {
		return nil, envErr
	}

	request.URL.Host = addr

	request.Header.Set("X-Request-Id", ctx.RequestID())
	request.Header.Set("X-User-Id", ctx.Actor().ID())

	return request, nil
}

func (s *GoodbyeService) responder(_ context.Context, _ *service.Request, response service.Response) service.Response {
	addCommonHeaders(response)

	return response
}
