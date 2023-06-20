package psql

import (
	"context"
	"time"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type messageRepository struct {
	db     *pgx.Conn
	logger *logrus.Entry
}

func (r *messageRepository) Create(ctx context.Context, m *models.Message) (*models.Message, store.StoreError) {
	q := `INSERT INTO messages(content, fk_user_id, fk_reply_message_id) VALUES ($1, $2, $3) RETURNING message_id, content, fk_user_id, fk_reply_message_id, created_at, modified_at;`

	r.logger.Tracef("SQL Query: %s", q)

	msg := new(models.Message)

	if err := r.db.QueryRow(ctx, q, m.Content, m.Username, m.Reply).Scan(&msg.Id, &msg.Content, &msg.Username, &msg.Reply, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return msg, nil
}

func (r *messageRepository) Find(ctx context.Context, id uint) (*models.Message, store.StoreError) {
	q := "SELECT message_id, fk_user_id, content, fk_reply_message_id, modified_at, created_at FROM messages WHERE message_id = $1 and is_deleted = false;"

	msg := new(models.Message)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, id).Scan(&msg.Id, &msg.Username, &msg.Content, &msg.Reply, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return msg, nil
}

func (r *messageRepository) FindAll(ctx context.Context, isIncludeDeleted bool) ([]*models.Message, store.StoreError) {
	q := "SELECT message_id, fk_user_id, content, fk_reply_message_id, modified_at, created_at FROM messages WHERE is_deleted = false ORDER BY message_id ASC;"
	if isIncludeDeleted {
		q = "SELECT message_id, fk_user_id, content, fk_reply_message_id, modified_at, created_at FROM messages ORDER BY message_id ASC;"
	}

	r.logger.Tracef("SQL Query: %s", q)

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	var msgs []*models.Message

	for rows.Next() {
		msg := new(models.Message)

		if err := rows.Scan(&msg.Id, &msg.Username, &msg.Content, &msg.Reply, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
			return nil, store.NewErrorAndLog(err, r.logger)
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func (r *messageRepository) EditWithUsername(ctx context.Context, content string, id uint, username string) (*models.Message, store.StoreError) {
	q := "UPDATE messages SET content = $1, modified_at = $2 WHERE message_id = $3 AND fk_user_id = $4 RETURNING message_id, content, fk_user_id, fk_reply_message_id, created_at, modified_at;"

	msg := new(models.Message)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, content, time.Now(), id, username).Scan(&msg.Id, &msg.Content, &msg.Username, &msg.Reply, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return msg, nil
}

func (r *messageRepository) Edit(ctx context.Context, content string, id uint) (*models.Message, store.StoreError) {
	q := "UPDATE messages SET content = $1, modified_at = $2 WHERE message_id = $3 RETURNING message_id, content, fk_user_id, fk_reply_message_id, created_at, modified_at;"

	msg := new(models.Message)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, content, time.Now(), id).Scan(&msg.Id, &msg.Content, &msg.Username, &msg.Reply, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return msg, nil
}

func (r *messageRepository) DeleteWithUsername(ctx context.Context, id uint, username string) store.StoreError {
	q := "UPDATE messages SET is_deleted = true WHERE message_id = $1 AND fk_user_id = $2;"

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, id, username); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *messageRepository) Delete(ctx context.Context, id uint) store.StoreError {
	q := "UPDATE messages SET is_deleted = true WHERE message_id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, id); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *messageRepository) DeletePermanently(ctx context.Context, id uint) store.StoreError {
	q := "DELETE FROM messages WHERE message_id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, id); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *messageRepository) Restore(ctx context.Context, id uint) (*models.Message, store.StoreError) {
	q := "UPDATE messages SET is_deleted = false WHERE message_id = $1 RETURNING message_id, content, fk_user_id, fk_reply_message_id, created_at, modified_at;"

	msg := new(models.Message)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, id).Scan(&msg.Id, &msg.Content, &msg.Username, &msg.Reply, &msg.CreatedAt, &msg.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return msg, nil
}
