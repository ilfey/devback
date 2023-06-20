package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/iservices"
	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
)

func Login(s *store.Store, jwt iservices.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := new(models.User)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "invalid body",
				"message": err.Error(),
			})
			return
		}

		user, err := s.User.Find(ctx.Request.Context(), body.Username)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "user error",
				"message": "user is not found",
			})
			return
		}

		if !user.ComparePassword(body.Password) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "user error",
				"message": "user is not found",
			})
			return
		}

		token := jwt.GenerateToken(user.Username)

		ctx.JSON(http.StatusOK, gin.H{
			"username":   user.Username,
			"created_at": user.CreatedAt,
			"token":      token,
		})
	}
}
