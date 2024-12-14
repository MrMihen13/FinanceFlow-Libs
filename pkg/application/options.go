package application

type Option func(*App) error

type ActionFunc func(*App) error

func emptyActionFunc(_ *App) error { return nil }

func (o Option) apply(app *App) error { return o(app) }
