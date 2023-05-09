package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
)

func CreateLink(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := new(models.Link)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "invalid body",
				"message": err.Error(),
			})
			return
		}

		if err := s.Link.Create(ctx.Request.Context(), body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "create link error",
				"message": "link not created",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
