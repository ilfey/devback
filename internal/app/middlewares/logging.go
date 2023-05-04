package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleWare(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logger.WithFields(logrus.Fields{
			"remote_addr": ctx.ClientIP(),
		})

		log.Infof("started %s %s", ctx.Request.Method, ctx.Request.RequestURI)

		start := time.Now()
		ctx.Next()

		code := ctx.Writer.Status()
		var level logrus.Level

		switch {
		case code >= 500:
			level = logrus.ErrorLevel
		case code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}

		log.Logf(
			level,
			"completed with %d %s in %v",
			code,
			http.StatusText(code),
			time.Since(start),
		)
	}

}
