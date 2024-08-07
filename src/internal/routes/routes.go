package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ziipojii/OpsAPI.git/docs" //swaggo

	"github.com/ziipojii/OpsAPI.git/src/internal/handlers"
	"github.com/ziipojii/OpsAPI.git/src/internal/infrastructure/config"
)

func DefaultRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Register handlers with config
	r.GET("/application/health", func(c *gin.Context) {
		handlers.HealthCheck(c, cfg)
	})
	r.GET("/", func(c *gin.Context) {
		handlers.HealthCheck(c, cfg)
	})
	r.GET("/ping", handlers.PingHandler)

	// Base path (/service_name/path)
	base := r.Group("/" + cfg.App.ServiceName)
	ServiceRoutes(base, cfg)

	// swaggo
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func ServiceRoutes(base *gin.RouterGroup, cfg *config.Config) {
	// base.GET("/ping", handlers.PingHandler)

	base.POST("/insertwfapi", func(c *gin.Context) {
		handlers.InsertWfapi(c, cfg)
	})
}
