package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func EditMessage(s *store.Store) gin.HandlerFunc {
	type request struct {
		Id      uint   `json:"id" binding:"required,min=1"`
		Content string `json:"content" binding:"required,min=1,max=2000"`
	}
	return func(ctx *gin.Context) {
		body := new(request)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "invalid body error",
				"message": err.Error(),
			})
			return
		}

		if err := s.Message.Edit(ctx.Request.Context(), body.Content, body.Id); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "edit message error",
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
