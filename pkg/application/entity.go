package application

type entity struct {
	name     string
	initFunc ActionFunc
	runFunc  ActionFunc
	stopFunc ActionFunc
}

type entities map[string]entity

func (e entities) add(ent entity) bool {
	if _, ok := e[ent.name]; ok {
		return false
	}

	e[ent.name] = ent

	return true
}

func (e entities) init(a *App) error {
	ctx := a.Ctx
	actionsList := make([]ActionFunc, 0, len(e))
	for _, el := range e {
		actionsList = append(actionsList, el.initFunc)
	}

	return a.runParallel(ctx, convertActionFuncToInterfaces(actionsList...)...)
}

func (e entities) run(a *App) error {
	ctx := a.Ctx
	actionsList := make([]ActionFunc, 0, len(e))

	for _, el := range e {
		actionsList = append(actionsList, el.runFunc)
	}
	return a.runParallel(ctx, convertActionFuncToInterfaces(actionsList...)...)
}

func (e entities) stop(a *App) error {
	ctx := a.Ctx

	actionsList := make([]ActionFunc, 0, len(e))
	for _, el := range e {
		actionsList = append(actionsList, el.stopFunc)
	}

	return a.runParallel(ctx, convertActionFuncToInterfaces(actionsList...)...)
}
