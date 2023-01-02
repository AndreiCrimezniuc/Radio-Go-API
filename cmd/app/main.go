package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nokogiriwatir/radio-main/internal/config"
	"nokogiriwatir/radio-main/internal/database"
	"nokogiriwatir/radio-main/internal/logger"
	"nokogiriwatir/radio-main/internal/router_handler"
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

	router.LoadHTMLGlob("assets/html/index.html")

	router.GET("/", env.HandleRoot)

	router.GET("/stations/:slug", env.HandleStation)

	router.Run(configs.App.AppPort)
}
