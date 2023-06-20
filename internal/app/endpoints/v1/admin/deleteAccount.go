package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func DeleteAccount(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")

		if err := s.User.Delete(ctx.Request.Context(), username); err != nil {
			switch err.Type() {
			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "account delete error",
					"message": "account not deleted",
				})
			}
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
