package config

import (
	"database/sql"
	"fmt"
)

var (
	db     *sql.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializePostgres()
	if err != nil {
		return fmt.Errorf("error initializing postgres: %v", err)
	}

	return nil
}

func GetPostgres() *sql.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
