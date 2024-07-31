package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ziipojii/OpsAPI.git/src/internal/infrastructure/config"
)

// InsertWfapi handles the insertion of JSON data into MongoDB
// @Summary Insert data into MongoDB
// @Description Insert JSON data into MongoDB with service_name, environment, and jenkins_host headers
// @Tags wfapi
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "JSON payload"
// @Param X-Service-Name header string true "Service Name"
// @Param X-Environment header string true "Environment"
// @Param X-Jenkins-Host header string true "Jenkins Host"
// @Success 201 {object} map[string]string "Successfully inserted"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tix-devops-api/insert [post]
func InsertWfapi(c *gin.Context, cfg *config.Config) {
	var result map[string]interface{}

	// Read JSON from request body
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Get headers from request
	serviceName := c.GetHeader("X-Service-Name")
	environment := c.GetHeader("X-Environment")
	jenkinsHost := c.GetHeader("X-Jenkins-Host")

	// Validate required headers
	if serviceName == "" || environment == "" || jenkinsHost == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required headers"})
		return
	}

	// Add headers to the result
	result["service_name"] = serviceName
	result["environment"] = environment
	result["jenkins_host"] = jenkinsHost

	// Connect to MongoDB
	dbClient, err := config.NewDatabases(cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	// Get the collection
	collection := dbClient.Database(cfg.Database.DBName).Collection("wfapi_results")

	// Insert the result into the collection
	if _, err := collection.InsertOne(context.Background(), result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert into MongoDB"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Inserted successfully"})
}
