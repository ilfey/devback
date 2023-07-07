package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		aCtx := ctx.MustGet(AUTH_CONTEXT).(*AuthorizationContext)

		// Check authorization
		if !aCtx.IsAuthorized() {
			// If not authorized return 404
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.Next()
	}
}
