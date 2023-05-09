package psql

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type linkRepository struct {
	db     *pgx.Conn
	logger *logrus.Entry
}

func (r *linkRepository) Create(ctx context.Context, m *models.Link) error {
	q := `INSERT INTO links(description, url) VALUES ($1, $2);`

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, m.Description, m.Url)
	if err != nil {
		r.logger.Errorf("Unknown create error: %v", err)

		return err
	}

	return nil
}

func (r *linkRepository) Find(ctx context.Context, id uint) (*models.Link, error) {
	q := `SELECT id, description, url, modified_at, created_at FROM links WHERE id = $1;`

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, id).Scan(&l.Id, &l.Description, &l.Url, &l.ModifiedAt, &l.CreatedAt); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			r.logger.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())

			return nil, err
		}

		r.logger.Errorf("Scan user in Find method error: %v", err)

		return nil, err
	}

	return l, nil
}

func (r *linkRepository) FindAll(ctx context.Context, isDeleted bool) ([]*models.Link, error) {
	q := `SELECT id, description, url, modified_at, created_at FROM links WHERE is_deleted = $1;`

	r.logger.Tracef("SQL Query: %s", q)

	rows, err := r.db.Query(ctx, q, isDeleted)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			r.logger.Errorf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())

			return nil, err
		}

		r.logger.Errorf("Unknown ReadAll error: %v", err)

		return nil, err
	}

	var links []*models.Link

	for rows.Next() {
		link := new(models.Link)

		err := rows.Scan(&link.Id, &link.Description, &link.Url, &link.ModifiedAt, &link.CreatedAt)
		if err != nil {
			r.logger.Errorf("Scan ReadAll error: %v", err)

			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}

func (r *linkRepository) Delete(ctx context.Context, id uint) error {
	q := "UPDATE links SET is_deleted = true WHERE id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		r.logger.Errorf("Unknown Delete error: %v", err)

		return err
	}

	return nil
}

func (r *linkRepository) DeletePermanently(ctx context.Context, id uint) error {
	q := "DELETE FROM links WHERE id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		r.logger.Errorf("Unknown DeletePermanently error: %v", err)

		return err
	}

	return nil
}

func (r *linkRepository) Edit(ctx context.Context, description, url string, id uint) error {
	q := "UPDATE links SET description = $1, url = $2 WHERE id = $3"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, description, url, id)
	if err != nil {
		r.logger.Errorf("Unknown Restore error: %v", err)

		return err
	}

	return nil
}

func (r *linkRepository) EditUrl(ctx context.Context, url string, id uint) error {
	q := "UPDATE links SET url = $1 WHERE id = $2"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, url, id)
	if err != nil {
		r.logger.Errorf("Unknown Restore error: %v", err)

		return err
	}

	return nil
}

func (r *linkRepository) EditDescription(ctx context.Context, description string, id uint) error {
	q := "UPDATE links SET description = $1 WHERE id = $2"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, description, id)
	if err != nil {
		r.logger.Errorf("Unknown Restore error: %v", err)

		return err
	}

	return nil
}

func (r *linkRepository) Restore(ctx context.Context, id uint) error {
	q := "UPDATE links SET is_deleted = false WHERE id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		r.logger.Errorf("Unknown Restore error: %v", err)

		return err
	}

	return nil
}
