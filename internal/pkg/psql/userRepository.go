package psql

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	generator *QueryGenerator
	db        *pgxpool.Pool
	logger    *logrus.Entry
}

func (r *userRepository) Create(ctx context.Context, user *models.User) (*models.User, store.StoreError) {
	q := r.generator.Insert([]string{"user_id", "password"})

	if err := user.BeforeCreate(); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	u := new(models.User)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, user.Username, user.Hash).Scan(&u.Username, &u.Hash, &u.IsDeleted, &u.CreatedAt, &u.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return u, nil
}

func (r *userRepository) Find(ctx context.Context, username string) (*models.User, store.StoreError) {
	q := r.generator.Select(SelectConfig{
		Attrs: []string{
			"user_id",
			"password",
			"created_at",
		},
		Condition: "user_id = $$ and is_deleted = false",
	})

	u := new(models.User)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, username).Scan(&u.Username, &u.Hash, &u.CreatedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return u, nil
}

func (r *userRepository) ResetPassword(ctx context.Context, user *models.User) (*models.User, store.StoreError) {
	q := r.generator.Update([]string{"password"}, "user_id = $$")

	if err := user.BeforeCreate(); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	u := new(models.User)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, user.Hash, user.Username).Scan(&u.Username, &u.Hash, &u.IsDeleted, &u.CreatedAt, &u.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return u, nil
}

func (r *userRepository) Delete(ctx context.Context, username string) store.StoreError {
	q := r.generator.Update([]string{"is_deleted"}, "user_id = $$")

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, true, username); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *userRepository) DeletePermanently(ctx context.Context, username string) store.StoreError {
	q := r.generator.Delete("user_id = $$")

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, username); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *userRepository) Restore(ctx context.Context, username string) (*models.User, store.StoreError) {
	q := r.generator.Update([]string{"is_deleted"}, "user_id = $$")

	u := new(models.User)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, false, username).Scan(&u.Username, &u.Hash, &u.IsDeleted, &u.CreatedAt, &u.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return u, nil
}
