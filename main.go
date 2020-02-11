package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rconway/godo/routes"
)

func main() {
	engine := gin.Default()
	root := &engine.RouterGroup

	root.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.Login(root)
	routes.Todo(root)

	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
