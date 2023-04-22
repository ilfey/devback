package psql

import (
	"context"

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

func (r *messageRepository) ReadAll(ctx context.Context) ([]*models.Message, error) {
	q := "SELECT id, userId, content, reply, modified_at, created_at FROM messages;"

	r.logger.Tracef("SQL Query: %s", q)

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			r.logger.Errorf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())

			return nil, err
		}

		r.logger.Errorf("Unknown ReadAll error: %v", err)

		return nil, err
	}

	var msgs []*models.Message

	for rows.Next() {
		msg := new(models.Message)

		err := rows.Scan(&msg.Id, &msg.Username, &msg.Content, &msg.Reply, &msg.CreatedAt, &msg.ModifiedAt)
		if err != nil {
			r.logger.Errorf("Scan ReadAll error: %v", err)

			return nil, err
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func (r *messageRepository) Edit()   {}
func (r *messageRepository) Delete() {}
