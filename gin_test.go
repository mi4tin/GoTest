package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestGin(t *testing.T){
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.Run(":8081")
}
