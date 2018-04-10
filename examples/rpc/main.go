package main

import (
	"context"
	"expvar"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ansel1/merry"
	consul "github.com/hashicorp/consul/api"
	"go.uber.org/zap"

	"github.com/percolate/shisa/examples/rpc/service"
	"github.com/percolate/shisa/httpx"
	"github.com/percolate/shisa/lb"
	"github.com/percolate/shisa/sd"
)

const (
	timeFormat = "2006-01-02T15:04:05+00:00"
	name       = "hello"
)

func main() {
	start := time.Now().UTC()

	startTime := new(expvar.String)
	startTime.Set(start.Format(timeFormat))

	helloVar := expvar.NewMap("hello")
	helloVar.Set("start-time", startTime)
	helloVar.Set("uptime", expvar.Func(func() interface{} {
		now := time.Now().UTC()
		return now.Sub(start).String()
	}))

	addr := flag.String("addr", ":0", "service address")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("initializing logger: %v", err)
	}

	defer logger.Sync()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	client, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		logger.Fatal("consul client failed to initialize", zap.Error(err))
	}

	reg := sd.NewConsul(client)
	bal := lb.NewLeastN(reg, 2)

	service := &hello.Hello{
		Balancer: bal,
		Logger:   logger,
	}
	rpc.Register(service)
	rpc.HandleHTTP()

	server := http.Server{}

	listener, err := httpx.HTTPListenerForAddress(*addr)
	if err != nil {
		logger.Fatal("opening listener", zap.Error(err))
	}
	logger.Info("starting hello service", zap.String("addr", listener.Addr().String()))

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Serve(listener)
	}()

	saddr := listener.Addr().String()
	ru := &url.URL{
		Host:     saddr,
		RawQuery: fmt.Sprintf("id=%s", saddr),
	}
	if err := reg.Register(name, ru); err != nil {
		logger.Fatal("service failed to register", zap.Error(err))
	}
	defer reg.Deregister(name)

	cu := &url.URL{
		Scheme:   "tcp",
		Host:     saddr,
		Path:     rpc.DefaultRPCPath,
		RawQuery: fmt.Sprintf("interval=30s&id=%s&serviceid=%s", saddr, saddr),
	}

	if err := reg.AddCheck(name, cu); err != nil {
		logger.Fatal("healthcheck failed to register", zap.Error(err))
	}
	defer reg.RemoveChecks(name)

	select {
	case err := <-errCh:
		if !merry.Is(err, http.ErrServerClosed) {
			logger.Error("starting server", zap.Error(err))
		}
	case <-interrupt:
		server.Shutdown(context.Background())
	}
}
