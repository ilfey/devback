package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func DeleteAccount(s *store.Store) gin.HandlerFunc {
	type req struct {
		Username    string `json:"username" binding:"required,alpha,min=3,max=16"`
		Permanently bool   `json:"permanently" binding:"required"`
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
			if err := s.User.DeletePermanently(ctx.Request.Context(), body.Username); err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "account delete error",
					"message": "account not deleted",
				})
				return
			}
		} else {
			if err := s.User.Delete(ctx.Request.Context(), body.Username); err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "account delete error",
					"message": "account not deleted",
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
