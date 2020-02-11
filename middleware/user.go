package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// EnsureUser zzz
func EnsureUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		username, err := c.Cookie("username")

		if err != nil {
			log.Println("...NO USERNAME - redirect to /login...")
			c.Redirect(303, "../auth/login")
		} else {
			c.Set("username", username)
			log.Printf("...username = %v...\n", username)
			c.Next()
		}
	}
}
