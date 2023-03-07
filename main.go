package main

import (
	"github.com/athanstan/goskraper/controllers"
	"github.com/athanstan/goskraper/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	route := gin.Default()

	route.GET("/ping", controllers.Index)

	route.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	route.Run() // listen and serve on 0.0.0.0:8080
}
