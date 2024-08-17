package router

import (
	"github.com/me2seeks/cola/internal/controllers"

	"github.com/gin-gonic/gin"
)

var publicController = new(controllers.PublicController)

func LoadPublicRoutes(r *gin.Engine) {
	public := r.Group("/public")
	{
		public.GET("/ping", publicController.Ping)
	}
}
