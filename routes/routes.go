package routes

import (
	"go-gin-fw/middleware"
	"go-gin-fw/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Auth routes
	r.GET("/login", service.ShowLogin)
	r.POST("/login", service.Login)
	r.GET("/register", service.ShowRegister)
	r.POST("/register", service.Register)
	r.GET("/logout", service.Logout)

	// Protected route
	auth := r.Group("/", middleware.AuthRequired())
	{
		auth.GET("/", service.Home)
		auth.GET("/users", service.ListUsers)
		auth.GET("/users/edit/:id", service.EditUserForm)
		auth.POST("/users/edit/:id", service.EditUser)
		auth.GET("/users/delete/:id", service.DeleteUserConfirm)
		auth.POST("/users/delete/:id", service.DeleteUser)
	}

}
