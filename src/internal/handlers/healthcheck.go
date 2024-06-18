package handlers

import (
	"time"
	"net/http"
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
func HealthCheck(c *gin.Context) {
	// Load application configuration
	cfg, err := config.Config()
	if err != nil {
        c.JSON(500, gin.H{
            "error": "Failed to load configuration",
        })
        return	
	}
	
	// JSON respond
	jsonResponse := gin.H{
		"message":   "Service " + cfg.ServiceName + " up and running",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	
	// send responsd JSON
	c.JSON(http.StatusOK, jsonResponse)

}
