package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func GetContacts(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		links, err := s.Contact.FindAll(ctx.Request.Context())
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "contacts read error",
				"message": "contacts not found",
			})
			return
		}

		if links == nil {
			ctx.JSON(http.StatusOK, make([]int, 0))
			return
		}

		ctx.JSON(http.StatusOK, links)
	}
}
