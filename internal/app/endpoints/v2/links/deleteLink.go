package links

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/middlewares"
	"github.com/ilfey/devback/internal/pkg/store"
)

func DeleteLink(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		permanently := false

		idString := ctx.Param("id")

		// Parse link id
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "error parse id",
				"message": err.Error(),
			})
			return
		}

		aCtx := ctx.MustGet(middlewares.AUTH_CONTEXT).(*middlewares.AuthorizationContext)

		// Get username
		username, ok := aCtx.IsAuthorizeRequired()
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "identification error",
				"message": "you are not identified",
			})
			return
		}

		// Check "include_deleted" param is exists
		_, ok = ctx.GetQuery("permanently")
		if ok {
			if aCtx.IsAdmin() { // Check user is admin
				permanently = true
			}
		}

		if permanently {
			// Delete permanently (for admin only)
			if err := s.Link.DeletePermanently(ctx.Request.Context(), uint(id)); err != nil {
				switch err.Type() {
				case store.StoreUnknown:
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error":   "link delete error",
						"message": "link not deleted",
					})
				}
				return
			}
		} else {
			// Delete
			if aCtx.IsAdmin() {
				// Delete
				if err := s.Link.Delete(ctx.Request.Context(), uint(id)); err != nil {
					switch err.Type() {
					case store.StoreUnknown:
						ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
							"error":   "link delete error",
							"message": "link not deleted",
						})
					}
					return
				}
			} else {
				// Delete with username
				if err := s.Link.DeleteWithUsername(ctx.Request.Context(), uint(id), username); err != nil {
					switch err.Type() {
					case store.StoreUnknown:
						ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
							"error":   "link delete error",
							"message": "link not deleted",
						})
					}
					return
				}
			}
		}

		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
