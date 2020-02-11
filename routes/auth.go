package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Login zzz
func Login(r *gin.RouterGroup) *gin.RouterGroup {
	auth := r.Group("auth")

	auth.GET("", func(c *gin.Context) {
		c.Redirect(307, "auth/check")
	})

	auth.GET("login", func(c *gin.Context) {
		c.SetCookie("username", "fred", 3600, "/", "localhost", false, true)
		c.Writer.Write([]byte("...LOGGED IN..."))
	})

	auth.GET("check", func(c *gin.Context) {
		username, err := c.Cookie("username")
		if err != nil {
			username = "<ERROR - could not get username>"
		}
		c.Writer.Write([]byte(fmt.Sprintf("...[auth].../check => Logged in as %v\n", username)))
	})

	auth.GET("logout", func(c *gin.Context) {
		c.SetCookie("username", "", -1, "/", "localhost", false, true)
		c.Writer.Write([]byte("...LOGGED OUT..."))
	})

	return auth
}
