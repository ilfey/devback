package psql

import (
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func NewStore(db *pgx.Conn, logger *logrus.Logger) *store.Store {
	s := new(store.Store)

	s.User = &userRepository{
		db: db,
		logger: logger.WithFields(logrus.Fields{
			"repository": "user",
		}),
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
