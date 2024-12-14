package application

import (
	"errors"
	"fmt"

	gogrpc "github.com/MrMihen13/FinanceFlow-Libs/pkg/grpc/server"
)

var ErrServiceEnabled = errors.New("service already enabled")

type Service entity

func NewService(name string, initFunc, runFunc ActionFunc) Service {
	return Service{
		name:     name,
		initFunc: initFunc,
		runFunc:  runFunc,
	}
}

func (e entities) addService(s Service) bool { return e.add(entity(s)) }

func WithServices(e ...Service) Option {
	return func(a *App) error {
		for _, el := range e {
			if added := a.services.addService(el); !added {
				return fmt.Errorf("%w: %s", ErrServiceEnabled, el.name)
			}
		}
		return nil
	}
}

var (
	GRPCServer = func(cfg *gogrpc.Config) Service {
		return NewService("grpc", emptyActionFunc, func(app *App) error {
			return startGRPCServer(app, cfg)
		})
	}
)
