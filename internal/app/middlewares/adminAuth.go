package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		aCtx := ctx.MustGet(AUTH_CONTEXT).(*AuthorizationContext)

		if !aCtx.IsAdmin() {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.Next()
	}
}
