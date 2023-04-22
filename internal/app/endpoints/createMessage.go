package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/iservices"
	"github.com/ilfey/devback/internal/pkg/store"
)

func CreateMessage(s *store.Store, jwt iservices.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := new(models.Message)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		username, err := jwt.GetTokenId(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		body.Username = username

		err = s.Message.Create(ctx.Request.Context(), body)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.AbortWithStatus(http.StatusOK)
	}
}
