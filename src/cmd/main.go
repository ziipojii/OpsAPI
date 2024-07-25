package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	_ "github.com/ziipojii/OpsAPI.git/docs" //swaggo
	"github.com/ziipojii/OpsAPI.git/src/internal/infrastructure/config"
	"github.com/ziipojii/OpsAPI.git/src/internal/routes"
)

func main() {
	// Load application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration")
	}

	// Set log level based on env
	if cfg.App.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(log.InfoLevel)
	} else {
		gin.SetMode(gin.DebugMode)
		log.SetLevel(log.DebugLevel)
	}

	// Set port apps
	port := os.Getenv("HTTP_SERVER_PORT")
	if port == "" {
		port = strconv.Itoa(cfg.App.HttpServerPort)
	}

	// Setup route with config
	r := routes.SetupRouter(cfg)

	r.Run(":" + port)
}
