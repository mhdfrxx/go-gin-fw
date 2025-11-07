package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	username, _ := c.Get("username")
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Username": username,
	})
}
