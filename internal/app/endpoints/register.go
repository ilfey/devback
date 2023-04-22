package endpoints

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
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := s.User.Create(ctx.Request.Context(), body); err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok {
				if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
					ctx.AbortWithStatus(http.StatusConflict)
				}
			}
			return
		}

		ctx.AbortWithStatus(http.StatusCreated)
	}
}
