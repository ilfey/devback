package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func ReadMessages(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		msgs, err := s.Message.ReadAll(ctx.Request.Context())
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, msgs)
	}
}
