package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/iservices"
)

func JwtAuthMiddleware(jwt iservices.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := jwt.GetToken(ctx)

		_, err := jwt.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "authorization error",
				"message": "not authorized",
			})
			return
		}

		ctx.Next()
	}
}
