package main

import (
    "os"
	"strconv"
    "github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

    "github.com/ziipojii/OpsAPI.git/src/internal/infrastructure/config"
	"github.com/ziipojii/OpsAPI.git/src/internal/routes"

	_ "github.com/ziipojii/OpsAPI.git/docs" //swaggo
)


func main()  {

	// Load application configuration
	cfg, err := config.Config()
	if err != nil {
		log.Fatal("Failed to load configuration")
	}

	// Set log level based on env
    if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
        log.SetLevel(log.InfoLevel)
    } else {
		gin.SetMode(gin.DebugMode)
        log.SetLevel(log.DebugLevel)
    }

	// Set port apps
	port := os.Getenv("HTTP_SERVER_PORT")
	if port == "" {
        port = strconv.Itoa(cfg.HttpServerPort)
	}

	// Setup route
	r := routes.SetupRouter()

	r.Run(":" + port) 
}
