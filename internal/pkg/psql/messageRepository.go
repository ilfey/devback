package psql

import (
	"context"
	"time"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type messageRepository struct {
	generator *QueryGenerator
	db        *pgxpool.Pool
	logger    *logrus.Entry
}

func (r *messageRepository) Create(ctx context.Context, m *models.Message) (*models.Message, store.StoreError) {
	q := r.generator.Insert([]string{
		"content",
		"fk_user_id",
		"fk_reply_message_id",
	})

	r.logger.Tracef("SQL Query: %s", q)

	msg := new(models.Message)

	if err := r.db.QueryRow(ctx, q, m.Content, m.Username, m.Reply).Scan(&msg.Id, &msg.Content, &msg.Username, &msg.Reply, &msg.IsDeleted, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return msg, nil
}

func (r *messageRepository) Find(ctx context.Context, id uint) (*models.Message, store.StoreError) {
	q := r.generator.Select(SelectConfig{
		Attrs: []string{
			"message_id",
			"fk_user_id",
			"content",
			"fk_reply_message_id",
			"is_deleted",
			"created_at",
			"modified_at",
		},
		Condition: "message_id = $$ and is_deleted = false",
	})

	msg := new(models.Message)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, id).Scan(&msg.Id, &msg.Username, &msg.Content, &msg.Reply, &msg.IsDeleted, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return msg, nil
}

func (r *messageRepository) FindAll(ctx context.Context, isIncludeDeleted bool) ([]*models.Message, store.StoreError) {
	config := SelectConfig{
		Attrs: []string{
			"message_id",
			"fk_user_id",
			"content",
			"fk_reply_message_id",
			"is_deleted",
			"created_at",
			"modified_at",
		},
		Condition: "is_deleted = false",
		OrderBy: []Order{
			{
				Attr: "message_id",
			},
		},
	}

	if isIncludeDeleted {
		config.Condition = ""
	}

	q := r.generator.Select(config)

	r.logger.Tracef("SQL Query: %s", q)

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	var msgs []*models.Message

	for rows.Next() {
		msg := new(models.Message)

		if err := rows.Scan(&msg.Id, &msg.Username, &msg.Content, &msg.Reply, &msg.IsDeleted, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
			return nil, store.NewErrorAndLog(err, r.logger)
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func (r *messageRepository) EditWithUsername(ctx context.Context, content string, id uint, username string) (*models.Message, store.StoreError) {
	q := r.generator.Update(
		[]string{
			"content",
			"modified_at",
		},
		"message_id = $$ and fk_user_id = $$",
	)

	msg := new(models.Message)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, content, time.Now(), id, username).Scan(&msg.Id, &msg.Content, &msg.Username, &msg.Reply, &msg.IsDeleted, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return msg, nil
}

func (r *messageRepository) Edit(ctx context.Context, content string, id uint) (*models.Message, store.StoreError) {
	q := r.generator.Update([]string{
		"content",
		"modified_at",
	},
		"message_id = $$",
	)

	msg := new(models.Message)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, content, time.Now(), id).Scan(&msg.Id, &msg.Content, &msg.Username, &msg.Reply, &msg.IsDeleted, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return msg, nil
}

func (r *messageRepository) DeleteWithUsername(ctx context.Context, id uint, username string) store.StoreError {
	q := r.generator.Update([]string{
		"is_deleted",
	},
		"message_id = $$ and fk_user_id = $$",
	)

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, true, id, username); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *messageRepository) Delete(ctx context.Context, id uint) store.StoreError {
	q := r.generator.Update([]string{
		"is_deleted",
	},
		"message_id = $$",
	)

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, true, id); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *messageRepository) DeletePermanently(ctx context.Context, id uint) store.StoreError {
	q := r.generator.Delete("message_id = $$")

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, id); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *messageRepository) Restore(ctx context.Context, id uint) (*models.Message, store.StoreError) {
	q := r.generator.Update(
		[]string{
			"is_deleted",
		},
		"message_id = $$",
	)

	msg := new(models.Message)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, false, id).Scan(&msg.Id, &msg.Content, &msg.Username, &msg.Reply, &msg.IsDeleted, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return msg, nil
}
