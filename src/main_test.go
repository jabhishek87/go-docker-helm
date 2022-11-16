package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPingRoute(t *testing.T) {
	r := SetUpRouter()
	r.GET("/", RoutePing)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)

}
