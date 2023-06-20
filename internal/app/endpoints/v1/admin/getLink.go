package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func GetLink(s *store.Store) gin.HandlerFunc {
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

		link, _err := s.Link.Find(ctx.Request.Context(), uint(id))
		if _err != nil {
			switch _err.Type() {
			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "internal server error",
					"message": "links read error",
				})

			case store.StoreNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{})
			}

			return
		}

		ctx.JSON(http.StatusOK, link)
	}
}
