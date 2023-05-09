package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func GetLinks(s *store.Store) gin.HandlerFunc {
	type request struct {
		IsDeleted bool `json:"is_deleted"`
	}
	return func(ctx *gin.Context) {
		body := new(request)

		if err := ctx.ShouldBindJSON(body); err != nil {
			body.IsDeleted = false
		}

		links, err := s.Link.FindAll(ctx.Request.Context(), body.IsDeleted)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "links read error",
				"message": "links not found",
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
