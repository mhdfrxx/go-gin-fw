package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Sessions = map[string]string{}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session_id")
		if err != nil || Sessions[cookie.Value] == "" {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		c.Set("username", Sessions[cookie.Value])
		c.Next()
	}
}
