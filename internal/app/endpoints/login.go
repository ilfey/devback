package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/iservices"
	"github.com/ilfey/devback/internal/pkg/store"
)

func Login(s *store.Store, jwt iservices.JWT) gin.HandlerFunc {
	type resp struct {
		Token string `json:"token"`
	}

	return func(ctx *gin.Context) {
		body := new(models.User)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		user, err := s.User.Find(ctx.Request.Context(), body.Username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !user.ComparePassword(body.Password) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		r := resp{
			Token: jwt.GenerateToken(user.Username),
		}

		ctx.JSON(http.StatusOK, r)
	}
}
