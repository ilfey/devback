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

		link, err := s.Link.Create(ctx.Request.Context(), body)
		if err != nil {
			switch err.Type() {
			case store.StoreAlreadyExists:
				ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
					"error":   "create link error",
					"message": "link already created",
				})
			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "create link error",
					"message": "link not created",
				})
			}
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK, link)
	}
}
