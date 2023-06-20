package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/iservices"
	"github.com/ilfey/devback/internal/pkg/store"
)

func DeleteAccount(s *store.Store, jwt iservices.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, err := jwt.GetTokenId(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "identification error",
				"message": "you are not identified",
			})
			return
		}

		if err := s.User.Delete(ctx.Request.Context(), username); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "account delete error",
				"message": "account not deleted",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
