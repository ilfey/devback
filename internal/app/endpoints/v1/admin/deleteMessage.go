package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func DeleteMessage(s *store.Store) gin.HandlerFunc {
	type req struct {
		Id          uint `json:"id" binding:"required,min=1"`
		Permanently bool `json:"permanently" binding:"required"`
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

		if body.Permanently {
			if err := s.Message.DeletePermanently(ctx.Request.Context(), body.Id); err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "message delete permanently error",
					"message": err.Error(),
				})
				return
			}
		} else {
			if err := s.Message.Delete(ctx.Request.Context(), body.Id); err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "message deletion error",
					"message": err.Error(),
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
