package links

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/middlewares"
	"github.com/ilfey/devback/internal/pkg/store"
)

func RestoreLink(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idString := ctx.Param("id")

		aCtx := ctx.MustGet(middlewares.AUTH_CONTEXT).(*middlewares.AuthorizationContext)

		if !aCtx.IsAdmin() {
			aCtx.AbortWithStatus(http.StatusForbidden)
			return
		}

		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "error parse id",
				"message": err.Error(),
			})
			return
		}

		link, _err := s.Link.Restore(ctx.Request.Context(), uint(id))
		if _err != nil {
			switch _err.Type() {
			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "internal server error",
					"message": "link restore error",
				})
			}
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK, link)
	}
}
