package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/iservices"
	"github.com/ilfey/devback/internal/pkg/store"
)

func EditMessage(s *store.Store, jwt iservices.JWT) gin.HandlerFunc {
	type request struct {
		Id      uint   `json:"id" binding:"required,min=1"`
		Content string `json:"content" binding:"required,min=1,max=2000"`
	}
	return func(ctx *gin.Context) {
		body := new(request)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "invalid body error",
				"message": err.Error(),
			})
			return
		}

		username, err := jwt.GetTokenId(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "identification error",
				"message": "you are not identified",
			})
			return
		}

		if err := s.Message.EditWithUsername(ctx.Request.Context(), body.Content, body.Id, username); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "edit message error",
				"message": "message not edited",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
