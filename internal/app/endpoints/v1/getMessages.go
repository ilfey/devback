package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/middlewares"
	"github.com/ilfey/devback/internal/pkg/store"
)

func GetMessages(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cursor := 0
		limit := 100
		includeDeleted := false

		// Get cursor query
		cursorString := ctx.Query("cursor")
		if cursorString != "" {

			// Try parse cursor
			temp, err := strconv.Atoi(cursorString)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "bad request error",
					"message": "failed to parse \"cursor\" query",
				})
				return
			}

			// If cursor < 1
			if temp < 0 {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "bad request error",
					"message": "\"cursor\" query cannot be less than 0",
				})
				return
			}

			// Set cursor
			cursor = temp
		}

		// Get limit query
		limitString := ctx.Query("limit")
		if limitString != "" {

			// Try parse limit
			temp, err := strconv.Atoi(limitString)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "bad request error",
					"message": "failed to parse \"limit\" query",
				})
				return
			}

			// If limit not in [1, 100]
			if temp < 1 || temp > 100 {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "bad request error",
					"message": "\"limit\" query must take the values [1, 100]",
				})
				return
			}

			// Set limit
			limit = temp
		}

		aCtx := ctx.MustGet(middlewares.AUTH_CONTEXT).(*middlewares.AuthorizationContext)

		if aCtx.IsAdmin() {
			// Check "include_deleted" param is exists
			_, ok := ctx.GetQuery("include_deleted")
			if ok {
				includeDeleted = true
			}
		}

		msgs, err := s.Message.FindAllWithScrolling(ctx.Request.Context(), cursor, limit, includeDeleted)
		if err != nil {
			switch err.Type() {
			case store.StoreUnknown:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "internal server error",
					"message": "messages read error",
				})
			}

			return
		}

		if msgs == nil {
			ctx.JSON(http.StatusOK, make([]int, 0))
			return
		}

		ctx.JSON(http.StatusOK, msgs)
	}
}
