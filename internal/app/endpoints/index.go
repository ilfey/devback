package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"github": "https://github.com/ilfey/devback",
		})
	}
}
