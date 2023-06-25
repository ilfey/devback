package psql

import (
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func NewStore(db *pgxpool.Pool, logger *logrus.Logger) *store.Store {
	s := new(store.Store)

	s.User = &userRepository{
		db: db,
		logger: logger.WithFields(logrus.Fields{
			"repository": "user",
		}),
		generator: NewQueryGenerator(
			"users",
			[]string{
				"user_id",
				"password",
				"is_deleted",
				"created_at",
				"modified_at",
			},
		),
	}

	s.Message = &messageRepository{
		db: db,
		logger: logger.WithFields(logrus.Fields{
			"repository": "message",
		}),
	}

	s.Link = &linkRepository{
		db: db,
		logger: logger.WithFields(logrus.Fields{
			"repository": "link",
		}),
	}

	s.Contact = &contactRepository{
		db: db,
		logger: logger.WithFields(logrus.Fields{
			"repository": "contact",
		}),
	}

	return s
}
