package application

import (
	"context"
	"fmt"
	"reflect"
	"runtime"

	"golang.org/x/sync/errgroup"
)

func convertActionFuncToInterfaces(fns ...ActionFunc) []interface{} {
	out := make([]interface{}, 0, len(fns))
	for _, fn := range fns {
		out = append(out, fn)
	}
	return out
}
func (a *App) runParallel(ctx context.Context, funcs ...interface{}) error {

	var g errgroup.Group
	for _, f := range funcs {
		f := f

		g.Go(func() (err error) {
			name := getFuncName(f)

			defer func() {
				if err != nil {
					a.Log.ErrorContext(ctx, "Error with running "+name, err)
				}
			}()
			switch t := f.(type) {
			case ActionFunc:
				return t(a)
			default:
				return fmt.Errorf("unkonwn sugnature with func %s", name)
			}

		})
	}
	return g.Wait()
}

func getFuncName(i interface{}) string {
	if val := reflect.ValueOf(i); val.Kind() == reflect.Func {
		if funcForPC := runtime.FuncForPC(val.Pointer()); funcForPC != nil {
			return funcForPC.Name()
		}
	}
	return "unknown"
}
