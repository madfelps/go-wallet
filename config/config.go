package config

import "errors"

var (
	logger *Logger
)

func Init() error {
	return errors.New("fake Error")
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
