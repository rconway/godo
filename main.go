package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rconway/godo/routes"
)

func main() {
	engine := gin.Default()
	root := &engine.RouterGroup

	routes.Root(root)
	routes.Login(root)
	routes.Todo(root)

	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
