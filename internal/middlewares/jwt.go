package middlewares

import (
	"net/http"
	"strings"

	"github.com/me2seeks/cola/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Jwt middleware
func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}
		parts := strings.Split(tokenStr, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}

		token := parts[1]
		claims, err := utils.ParseToken(token)
		if err != nil || claims == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}
		ctx.Set("ID", claims.Subject)

		ctx.Next()
	}
}
