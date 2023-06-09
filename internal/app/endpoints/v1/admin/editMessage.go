package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/store"
)

func EditMessage(s *store.Store) gin.HandlerFunc {
	type request struct {
		Content string `json:"content" binding:"required,min=1,max=2000"`
	}
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

		body := new(request)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "invalid body error",
				"message": err.Error(),
			})
			return
		}

		msg, _err := s.Message.Edit(ctx.Request.Context(), body.Content, uint(id))
		if _err != nil {
			switch _err.Type() {
			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "edit message error",
					"message": "message not edited",
				})
			}
			return
		}

		ctx.JSON(http.StatusOK, msg)
	}
}
