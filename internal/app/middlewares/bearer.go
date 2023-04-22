package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/services"
)

func JwtAuthMiddleware(jwt services.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := jwt.GetToken(ctx)

		_, err := jwt.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
