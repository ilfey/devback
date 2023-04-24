package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/config"
	"github.com/ilfey/devback/internal/pkg/iservices"
)

func AdminAuthMiddleware(jwt iservices.JWT, config *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := jwt.GetToken(ctx)

		_, err := jwt.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		var username string

		username, err = jwt.GetTokenId(ctx)
		if err != nil || username != config.AdminUsername {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.Next()
	}
}
