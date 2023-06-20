package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func GetMessages(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: add query param
		msgs, err := s.Message.FindAll(ctx.Request.Context(), false)
		if err != nil {
			switch err.Type() {
			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "internal server error",
					"message": "messages read error",
				})
			}

			return
		}

		if msgs == nil {
			ctx.JSON(http.StatusOK, make([]int, 0))
			return
		}

		ctx.JSON(http.StatusOK, msgs)
	}
}
