package routes

import (
	"go-gin-fw/middleware"
	"go-gin-fw/service"

	"github.com/gin-gonic/gin"
)

func Routes() {

	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	// Auth routes
	r.GET("/login", service.ShowLogin)
	r.POST("/login", service.Login)
	r.GET("/register", service.ShowRegister)
	r.POST("/register", service.Register)
	r.GET("/logout", service.Logout)

	// Protected route
	r.GET("/", middleware.AuthRequired(), service.Home)

	r.Run(":8080")

}
