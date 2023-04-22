package psql

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	db     *pgx.Conn
	logger *logrus.Entry
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	q := "INSERT INTO users (username, password) VALUES ($1, $2);"

	if err := user.BeforeCreate(); err != nil {
		r.logger.Errorf("Error BeforeCreate: %v", err)

		return err
	}

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, user.Username, user.Hash)
	if err != nil {
		r.logger.Errorf("Unknown create error: %v", err)

		return err
	}

	return nil
}

func (r *userRepository) Find(ctx context.Context, username string) (*models.User, error) {
	q := "SELECT username, password, created_at FROM users WHERE username = $1;"

	u := new(models.User)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, username).Scan(&u.Username, &u.Hash, &u.CreatedAt); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			r.logger.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())

			return nil, err
		}

		r.logger.Errorf("Scan user in Find method error: %v", err)

		return nil, err
	}

	return u, nil
}

func (r *userRepository) ResetPassword(ctx context.Context, user *models.User) error {
	q := "UPDATE users SET password = $1 WHERE username = $2;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, user.Password, user.Username)
	if err != nil {
		r.logger.Errorf("Unknown ResetPassword error: %v", err)

		return err
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, username string) error {
	q := "DELETE FROM users WHERE username = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, username)
	if err != nil {
		r.logger.Errorf("Unknown Delete error: %v", err)

		return err
	}

	return nil
}
