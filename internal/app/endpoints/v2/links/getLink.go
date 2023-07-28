package links

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/middlewares"
	"github.com/ilfey/devback/internal/pkg/store"
)

func GetLink(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		includeDeleted := false

		idString := ctx.Param("id")

		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "error parse id",
				"message": err.Error(),
			})

			return
		}

		aCtx := ctx.MustGet(middlewares.AUTH_CONTEXT).(*middlewares.AuthorizationContext)

		// Check "include_deleted" param is exists
		_, ok := ctx.GetQuery("include_deleted")
		if ok {
			if aCtx.IsAdmin() { // Check user is admin
				includeDeleted = true
			}
		}

		link, _err := s.Link.Find(ctx.Request.Context(), uint(id), includeDeleted)
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
