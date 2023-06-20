package psql

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	db     *pgx.Conn
	logger *logrus.Entry
}

func (r *userRepository) Create(ctx context.Context, user *models.User) (*models.User, store.StoreError) {
	q := "INSERT INTO users (user_id, password) VALUES ($1, $2) RETURNING user_id, password, created_at, modified_at;"

	if err := user.BeforeCreate(); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	u := new(models.User)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, user.Username, user.Hash).Scan(&u.Username, &u.Hash, &u.CreatedAt, &u.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return u, nil
}

func (r *userRepository) Find(ctx context.Context, username string) (*models.User, store.StoreError) {
	q := "SELECT user_id, password, created_at FROM users WHERE user_id = $1 and is_deleted = false;"

	u := new(models.User)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, username).Scan(&u.Username, &u.Hash, &u.CreatedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return u, nil
}

func (r *userRepository) ResetPassword(ctx context.Context, user *models.User) (*models.User, store.StoreError) {
	q := "UPDATE users SET password = $1 WHERE user_id = $2 RETURNING user_id, password, created_at, modified_at;"

	if err := user.BeforeCreate(); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	u := new(models.User)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, user.Hash, user.Username).Scan(&u.Username, &u.Hash, &u.CreatedAt, &u.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return u, nil
}

func (r *userRepository) Delete(ctx context.Context, username string) store.StoreError {
	q := "UPDATE users SET is_deleted = true WHERE user_id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, username); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *userRepository) DeletePermanently(ctx context.Context, username string) store.StoreError {
	q := "DELETE FROM users WHERE user_id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, username); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *userRepository) Restore(ctx context.Context, username string) (*models.User, store.StoreError) {
	q := "UPDATE users SET is_deleted = false WHERE user_id = $1 RETURNING user_id, password, created_at, modified_at;"

	u := new(models.User)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, username).Scan(&u.Username, &u.Hash, &u.CreatedAt, &u.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return u, nil
}
