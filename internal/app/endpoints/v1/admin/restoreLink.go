package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func RestoreLink(s *store.Store) gin.HandlerFunc {
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

		if err := s.Link.Restore(ctx.Request.Context(), uint(id)); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "link restore error",
				"message": "link not restored",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
