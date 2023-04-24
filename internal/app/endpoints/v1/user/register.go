package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func Register(s *store.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := new(models.User)

		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "invalid body",
				"message": err.Error(),
			})
			return
		}

		if err := s.User.Create(ctx.Request.Context(), body); err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok {
				if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
					ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
						"error":   "user create error",
						"message": "user already exists",
					})
					return
				}
				ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
					"error":   "user create error",
					"message": "user not created",
				})
			}
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "success",
		})
	}
}
