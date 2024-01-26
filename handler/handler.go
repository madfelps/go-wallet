package handler

import (
	"database/sql"

	"github.com/madfelps/go-wallet/config"
)

var (
	logger *config.Logger
	db     *sql.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetPostgres()
}
