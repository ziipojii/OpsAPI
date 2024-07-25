package handlers

import (
    "github.com/gin-gonic/gin"
	"net/http"
)

// PingHandler godoc
// @Summary      Responds with a pong message
// @Description  Responds with a JSON object containing a pong message
// @Tags         health
// @Accept       json
// @Produce      json
// @Router       /ping [get]
// @Router       /tix-devops-api/ping [get]
func PingHandler(c *gin.Context) {
	
	jsonResponse := gin.H{
		"message": "pong",
	}

    c.JSON(http.StatusOK, jsonResponse)
}

