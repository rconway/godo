package routes

import "github.com/gin-gonic/gin"

// Login zzz
func Login(r *gin.RouterGroup) *gin.RouterGroup {
	login := r.Group("login")

	login.GET("", func(c *gin.Context) {
		c.Writer.Write([]byte("...[login]...root"))
	})

	login.GET("check", func(c *gin.Context) {
		c.Writer.Write([]byte("...[login].../check"))
	})

	return login
}
