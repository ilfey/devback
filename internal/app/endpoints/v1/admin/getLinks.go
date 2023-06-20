package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func GetLinks(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		links, err := s.Link.FindAll(ctx.Request.Context(), false) // TODO: add query param
		if err != nil {
			switch err.Type() {
			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "internal server error",
					"message": "links read error",
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
