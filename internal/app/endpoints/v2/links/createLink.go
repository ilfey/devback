package links

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/middlewares"
	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
)

func CreateLink(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := new(models.Link)

		// Bind body
		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "invalid body",
				"message": err.Error(),
			})
			return
		}

		aCtx := ctx.MustGet(middlewares.AUTH_CONTEXT).(*middlewares.AuthorizationContext)

		// Get username
		username, ok := aCtx.IsAuthorizeRequired()
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "identification error",
				"message": "you are not identified",
			})
			return
		}

		body.Username = username

		// Create link
		link, err := s.Link.Create(ctx.Request.Context(), body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "create link error",
				"message": "link not created",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK, link)
	}
}
