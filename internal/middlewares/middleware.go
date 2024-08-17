package middlewares

import "github.com/gin-gonic/gin"

// loads all middlewares, ErrorHandle, Cors
func LoadMiddlewares(r *gin.Engine) *gin.Engine {
	r.Use(Cors())

	return r
}
