package psql

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type imageRepository struct {
	generator *QueryGenerator
	db        *pgxpool.Pool
	logger    *logrus.Entry
}
