package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func RestoreMessage(s *store.Store) gin.HandlerFunc {
	type req struct {
		Id uint `json:"id" binding:"required,min=1"`
	}
	return func(ctx *gin.Context) {
		body := new(req)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "invalid body",
				"message": err.Error(),
			})
			return
		}

		if err := s.Message.Restore(ctx.Request.Context(), body.Id); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "message restore error",
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
