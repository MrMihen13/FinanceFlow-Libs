package application

import (
	"context"
	"os"
	"syscall"

	gogrpc "github.com/MrMihen13/FinanceFlow-Libs/pkg/grpc/server"
	"github.com/MrMihen13/FinanceFlow-Libs/pkg/logs"
	"gorm.io/gorm"
)

type App struct {
	Ctx context.Context
	Log *logs.Logger

	DB *gorm.DB

	GRPCServer   *gogrpc.Server
	GRPCHandlers []GRPCHandler

	shutdown chan os.Signal
	started  chan struct{}
	closed   chan struct{}

	services entities
	storages entities
}

func NewAppWithContext(ctx context.Context, log *logs.Logger, opts ...Option) (*App, error) {
	app := &App{
		Ctx: ctx, Log: log,
		shutdown: make(chan os.Signal, 1),
		started:  make(chan struct{}, 1),
		closed:   make(chan struct{}),
		storages: make(entities),
		services: make(entities),
	}

	for _, opt := range opts {
		if err := opt.apply(app); err != nil {
			return nil, err
		}
	}

	if err := app.storages.init(app); err != nil {
		return nil, err
	}
	if err := app.services.init(app); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() {
	a.Log.Debug(a.Ctx, "Starting application")

	if len(a.services) == 0 {
		a.Log.Warn(a.Ctx, "No services to start")
	}

	if err := a.services.run(a); err != nil {
		a.Log.Fatal(a.Ctx, "run services", err)
		panic(err)
	}

	select {
	case <-a.closed:
		a.Log.Warn(a.Ctx, "application closed before started")
	case a.started <- struct{}{}:
		close(a.started)
		a.Log.Info(a.Ctx, "application started")
	}
	// wait for close
	<-a.closed
}

func (a *App) Shutdown() {
	a.shutdown <- syscall.SIGTERM

	// wait application to close
	<-a.closed
}

func (a *App) gracefulShutdown() {
	defer a.Log.Debug(a.Ctx, "graceful shutdown completed")

	sig := <-a.shutdown
	a.Log.Info(a.Ctx, "graceful shutdown signal received", sig)

	if err := a.services.stop(a); err != nil {
		a.Log.ErrorContext(a.Ctx, "graceful shutdown failed", err)
	}

	close(a.closed)
}
