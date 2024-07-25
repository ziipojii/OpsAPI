package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ziipojii/OpsAPI.git/src/internal/infrastructure/config"
)

// HealthCheck godoc
// @Summary Show the status of the service
// @Description get the status of the service
// @Tags health
// @Accept  json
// @Produce  json
// @Router /application/health [get]
func HealthCheck(c *gin.Context, cfg *config.Config) {
	if cfg == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Configuration not provided",
		})
		return
	}

	// JSON response
	jsonResponse := gin.H{
		"message":   "Service " + cfg.App.ServiceName + " up and running",
		"timestamp": time.Now().Format(time.RFC3339),
	}

	// Send response JSON
	c.JSON(http.StatusOK, jsonResponse)
}
