package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ziipojii/OpsAPI.git/src/internal/infrastructure/config"
)

// MockConfig is a mock configuration for testing
func MockConfig() *config.Config {
	return &config.Config{
		App: &config.AppCfg{
			ServiceName: "tix-devops-api",
		},
		Database: &config.DatabaseCfg{},
	}
}

// TestHealthCheck tests the HealthCheck handler
func TestHealthCheck(t *testing.T) {
	// Set gin mode to test
	gin.SetMode(gin.TestMode)

	// Create a new gin router and register the handler
	r := gin.Default()
	cfg := MockConfig()
	r.GET("/application/health", func(c *gin.Context) {
		HealthCheck(c, cfg)
	})

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodGet, "/application/health", nil)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert response body
	responseBody := w.Body.String()
	assert.Contains(t, responseBody, `"message":"Service tix-devops-api up and running"`)
	assert.Contains(t, responseBody, `"timestamp"`)

	// Optionally, assert timestamp is within reasonable range
}
