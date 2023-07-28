package links

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/middlewares"
	"github.com/ilfey/devback/internal/pkg/store"
)

func GetLinks(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		includeDeleted := false

		aCtx := ctx.MustGet(middlewares.AUTH_CONTEXT).(*middlewares.AuthorizationContext)

		// Check "include_deleted" param is exists
		_, ok := ctx.GetQuery("include_deleted")
		if ok {
			if aCtx.IsAdmin() { // Check user is admin
				includeDeleted = true
			}
		}

		links, err := s.Link.FindAll(ctx.Request.Context(), includeDeleted)
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
