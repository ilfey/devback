package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func CreateContact(s *store.Store) gin.HandlerFunc {
	type request struct {
		Title  string `json:"title" binding:"required"`
		LinkId uint   `json:"link_id" binding:"required,gte=1"`
	}
	return func(ctx *gin.Context) {
		body := new(request)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "invalid body",
				"message": err.Error(),
			})
			return
		}

		contact, err := s.Contact.Create(ctx.Request.Context(), body.Title, body.LinkId)
		if err != nil {
			if err.Type() == store.StoreUnknown {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "create contact error",
					"message": "contact not created",
				})
				return
			}

			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "create contact error",
				"message": "contact not created",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK, contact)
	}
}
