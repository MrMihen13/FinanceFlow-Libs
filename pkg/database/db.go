package database

import (
	"context"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(_ context.Context, cfg *ConnConfig) (*gorm.DB, error) {
	connString, err := buildConnectionString(cfg)
	if err != nil {
		return nil, err
	}

	client, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func Close(client *gorm.DB) error {
	ormPool, err := client.DB()
	if err != nil {
		return err
	}

	return ormPool.Close()
}

func Ping(ctx context.Context, client *gorm.DB) error {
	db, err := client.DB()
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	if inErr := errors.Unwrap(err); inErr != nil {
		return inErr
	}

	return nil
}
