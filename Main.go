package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var (
		host   string = "0.0.0.0"
		port   int    = 8888
		router *gin.Engine
	)

	router = gin.New()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	router.Static("/images", "/images")

	router.Run(fmt.Sprintf("%s:%d", host, port))
}
