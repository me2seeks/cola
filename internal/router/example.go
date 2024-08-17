package router

import (
	"github.com/me2seeks/cola/internal/controllers"
	"github.com/me2seeks/cola/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var exampleController = new(controllers.ExampleController)

func LoadExampleRoutes(r *gin.Engine) {
	example := r.Group("/examples")
	example.Use(middlewares.Jwt())

	{
		example.POST("/createExample", exampleController.CreateExample)
		example.GET("/getExample", exampleController.GetExample)
		example.POST("/updateExample", exampleController.UpdateExample)
		example.POST("/deleteExample", exampleController.DeleteExample)
	}
}
