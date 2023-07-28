package psql

import (
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func NewStore(db *pgxpool.Pool, logger *logrus.Logger) *store.Store {
	s := new(store.Store)

	s.Contact = &contactRepository{
		db: db,
		logger: logger.WithFields(logrus.Fields{
			"repository": "contact",
		}),
	}

	s.Image = &imageRepository{
		db: db,
		logger: logger.WithFields(logrus.Fields{
			"repository": "image",
		}),
		generator: NewQueryGenerator(
			"images",
			[]string{
				"image_id",
				"fk_link_id",
				"is_deleted",
				"created_at",
				"modified_at",
			},
		),
	}

	s.Link = &linkRepository{
		db: db,
		logger: logger.WithFields(logrus.Fields{
			"repository": "link",
		}),
		generator: NewQueryGenerator(
			"links",
			[]string{
				"link_id",
				"fk_user_id",
				"description",
				"url",
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
		generator: NewQueryGenerator(
			"messages",
			[]string{
				"message_id",
				"content",
				"fk_user_id",
				"fk_reply_message_id",
				"is_deleted",
				"created_at",
				"modified_at",
			},
		),
	}

	s.Project = &projectRepository{
		db: db,
		logger: logger.WithFields(logrus.Fields{
			"repository": "project",
		}),
		generator: NewQueryGenerator(
			"projects",
			[]string{
				"project_id",
				"title",
				"description",
				"fk_source_link_id",
				"fk_url_link_id",
				"is_deleted",
				"modified_at",
				"created_at",
			},
		),
	}

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

	return s
}
