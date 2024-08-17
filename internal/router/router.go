package router

import (
	"github.com/gin-gonic/gin"
	"github.com/me2seeks/cola/internal/middlewares"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	// Global
	Router.Use(middlewares.Cors())

	// public routes, no auth required
	LoadPublicRoutes(Router)

	LoadExampleRoutes(Router)

	LoadUserRoutes(Router)

	// init swagger
	// Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
