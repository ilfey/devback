package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func RestoreAccount(s *store.Store) gin.HandlerFunc {
	type req struct {
		Username string `json:"username" binding:"required,alpha,min=3,max=16"`
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

		if err := s.User.Restore(ctx.Request.Context(), body.Username); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "user restore error",
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
