package server

import (
	"context"
	"fmt"
	"net"

	"github.com/MrMihen13/FinanceFlow-Libs/pkg/logs"
	"google.golang.org/grpc"
)

type Server struct {
	log  *logs.Logger
	GRPC *grpc.Server
	cfg  *Config
}

func NewServer(log *logs.Logger, cfg *Config) *Server {
	gRPCServer := grpc.NewServer()

	return &Server{
		log:  log,
		GRPC: gRPCServer,
		cfg:  cfg,
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return err
	}

	s.log.Info(context.Background(), "GRPC server started", logs.String("address", l.Addr().String()))

	if err := s.GRPC.Serve(l); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	defer s.log.Info(context.Background(), "GRPC server stopped")
	s.GRPC.GracefulStop()
}
