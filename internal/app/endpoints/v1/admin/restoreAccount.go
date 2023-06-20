package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func RestoreAccount(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		user, err := s.User.Restore(ctx.Request.Context(), username)
		if err != nil {
			switch err.Type() {
			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "internal server error",
					"message": "user restore failed",
				})
			}
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"username":   user.Username,
			"created_at": user.CreatedAt,
		})
	}
}
