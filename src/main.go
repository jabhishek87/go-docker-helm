package main

import (
	"fmt"
	"net/http"
	"strings"

	docs "go-docker/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db = make(map[string]string)

func main() {

	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/api/v1")
	// docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "Gin Swagger Example API"

	// Auth Group
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":   "bar",   // user:foo password:bar
		"manu":  "123",   // user:manu password:123
		"admin": "admin", // user:admin password:admin
	}))

	v1.GET("/ping", RoutePing)
	v1.GET("/user/:name", getParams)
	authorized.POST("admin", postParams)

	router.NoRoute(func(c *gin.Context) {
		// c.AbortWithStatus(http.StatusNotFound)
		c.JSON(404, gin.H{"code": http.StatusNotFound, "message": "Page not found"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} json "{"message":"pong"}"
// @Router /api/v1/ping [get]
func RoutePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GetParamsExample godoc
// @Summary get params example
// @Schemes
// @Description do get
// @Tags getParams
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/user/{name} [get]
// @Param name path string true "name"
func getParams(c *gin.Context) {
	user := c.Params.ByName("name")
	status, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": status})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "idle"})
	}
}

// PostParamsExample godoc
// @Summary post params example
// @Schemes
// @Description do post
// @Tags postParams
// @Accept json
// @Produce json
// @Success 200
// @Router /admin [post]
// @Param data body string true "{status:val, name:val}"
// @Param authorization header string true " Basic Zm9vOmJhcg=="
func postParams(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	var json struct {
		Status string `json:"status" binding:"required"`
		Name   string `json:"name" binding:"required"`
	}

	if c.Bind(&json) == nil {
		db[strings.ToLower(json.Name)] = strings.ToLower(json.Status)
		fmt.Println(user, json.Status, json.Name)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
