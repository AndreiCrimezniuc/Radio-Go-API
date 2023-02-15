package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nokogiriwatir/radio-main/internal/router_handler"
	"nokogiriwatir/radio-main/pkg/config"
	"nokogiriwatir/radio-main/pkg/database"
	"nokogiriwatir/radio-main/pkg/logger"
)

func main() {
	configs := config.InitConfig()
	logger := logger.GetLogger(configs.Logger)

	handleRouting(configs, logger)
}

func handleRouting(configs config.Configs, logger *zap.Logger) {
	router := gin.Default()
	router.Use(cors.Default())

	db := database.Connection(configs.Db)

	env := router_handler.DIEnv{
		Db:     db,
		Config: &configs,
		Logger: logger,
	}

	router.GET("/ru/stations/:slug", env.HandleStation)

	router.GET("/ru/stations", env.HandleStations)

	router.Run(configs.App.AppPort)
}
