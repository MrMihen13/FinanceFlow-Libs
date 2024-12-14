package database

import (
	"errors"
	"fmt"
)

func buildConnectionString(config *ConnConfig) (string, error) {
	if config == nil {
		return "", errors.New("config is nil")
	}

	sslMode := "disable"
	if config.SSLMode {
		sslMode = "require"
	}

	if config.Host == "" {
		return "", errors.New("host is required")
	}

	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Database, config.Password, sslMode,
	), nil
}
