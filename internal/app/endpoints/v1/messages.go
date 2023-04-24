package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func ReadMessages(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		msgs, err := s.Message.ReadAll(ctx.Request.Context())
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "messages read error",
				"message": "messages not found",
			})
			return
		}

		if msgs == nil {
			ctx.JSON(http.StatusOK, make([]int, 0))
			return
		}

		ctx.JSON(http.StatusOK, msgs)
	}
}
