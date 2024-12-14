package application

import (
	gogrpc "github.com/MrMihen13/FinanceFlow-Libs/pkg/grpc/server"
	"google.golang.org/grpc"
)

type GRPCHandler interface {
	RegisterGRPC(*grpc.Server)
}

func startGRPCServer(a *App, cfg *gogrpc.Config) error {
	a.GRPCServer = gogrpc.NewServer(a.Log, cfg)

	for _, h := range a.GRPCHandlers {
		h.RegisterGRPC(a.GRPCServer.GRPC)
	}

	if err := a.GRPCServer.Run(); err != nil {
		return err
	}

	return nil
}

func (a *App) RegisterGRPCHandlers(handlers ...GRPCHandler) {
	a.GRPCHandlers = append(a.GRPCHandlers, handlers...)
}
