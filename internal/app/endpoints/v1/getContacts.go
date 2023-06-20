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
			switch err.Type() {
			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "internal server error",
					"message": "contacts read error",
				})
			}

			return
		}

		if links == nil {
			ctx.JSON(http.StatusOK, make([]int, 0))
			return
		}

		ctx.JSON(http.StatusOK, links)
	}
}
