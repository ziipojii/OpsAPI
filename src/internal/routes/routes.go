package routes

import (
    "github.com/gin-gonic/gin"
	"github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
	_ "github.com/ziipojii/OpsAPI.git/docs" //swaggo

    "github.com/ziipojii/OpsAPI.git/src/internal/handlers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

	// healthcheck
	r.GET("/application/health", handlers.HealthCheck)
	r.GET("/", handlers.HealthCheck)


	r.GET("/ping", handlers.PingHandler)

	// swaggo
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}
