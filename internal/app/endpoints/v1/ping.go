package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/config"
	"github.com/ilfey/devback/internal/pkg/iservices"
)

func Ping(c *config.Config, jwt iservices.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var username, message string

		code := http.StatusOK
		token := jwt.GetToken(ctx)
		uptime := time.Since(c.StartTime)

		if token != "" {
			_, err := jwt.ValidateToken(token)
			if err != nil {
				code = http.StatusUnauthorized
				message = "token is invalid"
			} else {
				username, err = jwt.GetTokenId(ctx)
				if err != nil {
					code = http.StatusUnauthorized
					message = "you are not identified"
				}
			}
		} else {
			code = http.StatusUnauthorized
			message = "token is empty"
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":       code,
			"message":    message,
			"username":   username,
			"start_time": c.StartTime,
			"uptime":     uptime.String(),
		})
	}
}
