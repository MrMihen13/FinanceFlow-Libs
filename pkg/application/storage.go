package application

import (
	"errors"
	"fmt"

	"github.com/MrMihen13/FinanceFlow-Libs/pkg/database"
	"gorm.io/gorm"
)

var ErrStorageEnabled = errors.New("storage already enabled")

type storage entity

func newStorage(name string, initFunc ActionFunc) storage {
	return storage{
		name:     name,
		initFunc: initFunc,
		runFunc:  emptyActionFunc,
	}
}

func (e entities) addStorage(s storage) bool { return e.add(entity(s)) }

func WithStorage(e ...storage) Option {
	return func(a *App) error {
		for _, el := range e {
			if added := a.storages.addStorage(el); !added {
				return fmt.Errorf("%w: %s", ErrStorageEnabled, el.name)
			}
		}
		return nil
	}
}

var Gorm = func(connCfg *database.ConnConfig, cfg *gorm.Config) storage {
	return newStorage("database", func(a *App) error {
		return initGorm(a, connCfg, cfg)
	})
}
