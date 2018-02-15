package main

import (
	"os"
	"time"

	"github.com/ansel1/merry"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/percolate/shisa/authn"
	"github.com/percolate/shisa/auxiliary"
	"github.com/percolate/shisa/env"
	"github.com/percolate/shisa/gateway"
	"github.com/percolate/shisa/middleware"
	"github.com/percolate/shisa/service"
)

func serve(logger *zap.Logger, addr, debugAddr, healthcheckAddr string) {
	idp := &ExampleIdentityProvider{Env: env.DefaultProvider}

	authenticator, err := authn.NewBasicAuthenticator(idp, "example")
	if err != nil {
		logger.Fatal("creating authenticator", zap.Error(err))
	}
	authN := &middleware.Authentication{Authenticator: authenticator}

	gw := &gateway.Gateway{
		Address:         addr,
		HandleInterrupt: true,
		GracePeriod:     2 * time.Second,
		Handlers:        []service.Handler{authN.Service},
		Logger:          logger,
	}

	authZ := SimpleAuthorization{[]string{"user:1"}}
	debug := &auxiliary.DebugServer{
		HTTPServer: auxiliary.HTTPServer{
			Addr:           debugAddr,
			Authentication: authN,
			Authorizer:     authZ,
		},
		Logger: logger,
	}

	hello := NewHelloService(env.DefaultProvider)
	goodbye := NewGoodbyeService(env.DefaultProvider)

	healthcheck := &auxiliary.HealthcheckServer{
		HTTPServer: auxiliary.HTTPServer{
			Addr:           healthcheckAddr,
			Authentication: authN,
			Authorizer:     authZ,
		},
		Checkers: []auxiliary.Healthchecker{idp, hello, goodbye},
		Logger:   logger,
	}

	services := []service.Service{hello, goodbye}

	if err := gw.Serve(services, debug, healthcheck); err != nil {
		for _, e := range multierr.Errors(err) {
			values := merry.Values(e)
			fs := make([]zapcore.Field, 0, len(values))
			for name, value := range values {
				if key, ok := name.(string); ok {
					fs = append(fs, zap.Reflect(key, value))
				}
			}
			logger.Error(merry.Message(e), fs...)
		}
		os.Exit(1)
	}
}
