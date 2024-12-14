package application

import "github.com/MrMihen13/FinanceFlow-Libs/pkg/database"

func initGorm(a *App, cfg *database.ConnConfig) error {
	var err error
	a.DB, err = database.Connect(a.Ctx, cfg)
	if err != nil {
		a.Log.Fatal(a.Ctx, "Failed to connect to database", err)
		return err
	}

	if err := database.Ping(a.Ctx, a.DB); err != nil {
		a.Log.Fatal(a.Ctx, "Failed to ping database", err)
	}

	a.Log.Debug(a.Ctx, "Connected to database")
	return nil
}
