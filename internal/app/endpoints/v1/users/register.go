package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
)

func Register(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := new(models.User)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "invalid body",
				"message": err.Error(),
			})
			return
		}

		user, err := s.User.Create(ctx.Request.Context(), body)

		if err != nil {
			switch err.Type() {
			case store.StoreAlreadyExists:
				ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
					"error":   "user create error",
					"message": "user already exists",
				})

			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "user create error",
					"message": "user not created",
				})
			}
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"username":   user.Username,
			"created_at": user.CreatedAt,
		})
	}
}
