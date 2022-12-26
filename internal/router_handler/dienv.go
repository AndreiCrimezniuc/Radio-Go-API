package router_handler

import (
	"database/sql"
	"go.uber.org/zap"
	"nokogiriwatir/radio-main/internal/config"
)

// Dependency Injection Environment

type DIEnv struct {
	Db     *sql.DB
	Config *config.Configs
	Logger *zap.Logger
}
