package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	// Set gin mode to test
	gin.SetMode(gin.TestMode)

	// Create a new gin router and register the handler
	r := gin.Default()
	r.GET("/ping", PingHandler)

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert response body
	expectedResponse := `{"message":"pong"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
