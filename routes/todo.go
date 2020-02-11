package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rconway/godo/middleware"
)

// Todo zzz
func Todo(r *gin.RouterGroup) *gin.RouterGroup {
	login := r.Group("todo")

	login.Use(middleware.EnsureUser())

	login.GET("", func(c *gin.Context) {
		username, usernameExists := c.Get("username")
		if !usernameExists {
			log.Println("ERROR getting username")
			username = "<ERROR getting username>"
		}
		c.Writer.Write([]byte(fmt.Sprintf("...[todo]...root with user = %v", username)))
	})

	return login
}
