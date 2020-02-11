package routes

import (
	"github.com/gin-gonic/gin"
)

// Login zzz
func Login(r *gin.RouterGroup) *gin.RouterGroup {
	login := r.Group("login")

	login.GET("", func(c *gin.Context) {
		c.SetCookie("username", "fred", 3600, "/", "localhost", false, true)
		c.Writer.Write([]byte("...LOGGED IN..."))
	})

	login.GET("check", func(c *gin.Context) {
		c.Writer.Write([]byte("...[login].../check"))
	})

	login.GET("logout", func(c *gin.Context) {
		c.SetCookie("username", "", -1, "/", "localhost", false, true)
		c.Writer.Write([]byte("...LOGGED OUT..."))
	})

	return login
}
