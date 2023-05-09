package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func GetContact(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idString := ctx.Param("id")

		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "error parse id",
				"message": err.Error(),
			})
			return
		}

		link, err := s.Contact.Find(ctx.Request.Context(), uint(id))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "contact read error",
				"message": "contact not found",
			})
			return
		}

		ctx.JSON(http.StatusOK, link)
	}
}
