package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func RestoreAccount(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")

		if err := s.User.Restore(ctx.Request.Context(), username); err != nil {
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
