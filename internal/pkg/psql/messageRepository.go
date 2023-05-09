package psql

import (
	"context"
	"time"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type messageRepository struct {
	db     *pgx.Conn
	logger *logrus.Entry
}

func (r *messageRepository) Create(ctx context.Context, m *models.Message) error {
	q := `INSERT INTO messages(content, userId, reply) VALUES ($1, $2, $3);`

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, m.Content, m.Username, m.Reply)
	if err != nil {
		r.logger.Errorf("Unknown create error: %v", err)

		return err
	}

	return nil
}

func (r *messageRepository) FindAll(ctx context.Context) ([]*models.Message, error) {
	q := "SELECT id, userId, content, reply, modified_at, created_at FROM messages WHERE is_deleted = false ORDER BY id ASC;"

	r.logger.Tracef("SQL Query: %s", q)

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			r.logger.Errorf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())

			return nil, err
		}

		r.logger.Errorf("Unknown FindAll error: %v", err)

		return nil, err
	}

	var msgs []*models.Message

	for rows.Next() {
		msg := new(models.Message)

		err := rows.Scan(&msg.Id, &msg.Username, &msg.Content, &msg.Reply, &msg.CreatedAt, &msg.ModifiedAt)
		if err != nil {
			r.logger.Errorf("Scan FindAll error: %v", err)

			return nil, err
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func (r *messageRepository) EditWithUsername(ctx context.Context, content string, id uint, username string) error {
	q := "UPDATE messages SET content = $1, modified_at = $2 WHERE id = $3 AND userId = $4;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, content, time.Now(), id, username)
	if err != nil {
		r.logger.Errorf("Unknown EditWithUsername error: %v", err)

		return err
	}

	return nil
}

func (r *messageRepository) Edit(ctx context.Context, content string, id uint) error {
	q := "UPDATE messages SET content = $1, modified_at = $2 WHERE id = $3;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, content, time.Now(), id)
	if err != nil {
		r.logger.Errorf("Unknown Edit error: %v", err)

		return err
	}

	return nil
}

func (r *messageRepository) DeleteWithUsername(ctx context.Context, id uint, username string) error {
	q := "DELETE FROM messages WHERE id = $1 AND userId = $2;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, id, username)
	if err != nil {
		r.logger.Errorf("Unknown DeleteWithUsername error: %v", err)

		return err
	}

	return nil
}

func (r *messageRepository) Delete(ctx context.Context, id uint) error {
	q := "UPDATE messages SET is_delete = true WHERE id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		r.logger.Errorf("Unknown Delete error: %v", err)

		return err
	}

	return nil
}

func (r *messageRepository) DeletePermanently(ctx context.Context, id uint) error {
	q := "DELETE FROM messages WHERE id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		r.logger.Errorf("Unknown Delete error: %v", err)

		return err
	}

	return nil
}

func (r *messageRepository) Restore(ctx context.Context, id uint) error {
	q := "UPDATE messages SET is_delete = false WHERE id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		r.logger.Errorf("Unknown Delete error: %v", err)

		return err
	}

	return nil
}
