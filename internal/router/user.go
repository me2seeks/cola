package router

import (
	"github.com/gin-gonic/gin"
	"github.com/me2seeks/cola/internal/controllers"
	"github.com/me2seeks/cola/internal/middlewares"
)

func LoadUserRoutes(r *gin.Engine) {
	userController := controllers.NewUserController()

	user := r.Group("user")

	{
		user.POST("/login", userController.Login)
		user.POST("/register", userController.Register)
	}

	{
		user.Use(middlewares.Jwt())
		user.GET("", userController.GetUser)
		user.PUT("", userController.UpdateUser)
		user.DELETE("", userController.DeleteUser)
	}

	users := r.Group("users")
	{
		users.Use(middlewares.Jwt())
		users.GET("", userController.ListUser)
	}
}
