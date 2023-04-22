package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/config"
	"github.com/ilfey/devback/internal/pkg/services"
)

func AdminAuthMiddleware(jwt services.JWT, config *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := jwt.GetToken(ctx)

		_, err := jwt.ValidateToken(token)
		if err != nil {
			print("token invalid")
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		var username string
		username, err = jwt.GetTokenId(ctx)
		print(username)
		print(config.AdminUsername)
		if err != nil || username != config.AdminUsername {
			print("username")
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.Next()
	}
}
