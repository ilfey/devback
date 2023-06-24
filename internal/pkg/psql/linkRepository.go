package psql

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type linkRepository struct {
	db     *pgxpool.Pool
	logger *logrus.Entry
}

func (r *linkRepository) Create(ctx context.Context, m *models.Link) (*models.Link, store.StoreError) {
	q := `INSERT INTO links(description, url) VALUES ($1, $2) RETURNING link_id, url, description, created_at, modified_at;`

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, m.Description, m.Url).Scan(&l.Id, &l.Url, &l.Description, &l.CreatedAt, &l.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}

func (r *linkRepository) Find(ctx context.Context, id uint) (*models.Link, store.StoreError) {
	q := `SELECT link_id, description, url, modified_at, created_at FROM links WHERE link_id = $1;`

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, id).Scan(&l.Id, &l.Description, &l.Url, &l.ModifiedAt, &l.CreatedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}

func (r *linkRepository) FindAll(ctx context.Context, isIncludeDeleted bool) ([]*models.Link, store.StoreError) {
	q := `SELECT link_id, description, url, modified_at, created_at FROM links WHERE is_deleted = $1;`

	if isIncludeDeleted {
		q = `SELECT link_id, description, url, modified_at, created_at FROM links;`
	}

	r.logger.Tracef("SQL Query: %s", q)

	rows, err := r.db.Query(ctx, q, isIncludeDeleted)
	if err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	var links []*models.Link

	for rows.Next() {
		link := new(models.Link)

		if err := rows.Scan(&link.Id, &link.Description, &link.Url, &link.ModifiedAt, &link.CreatedAt); err != nil {
			return nil, store.NewErrorAndLog(err, r.logger)
		}

		links = append(links, link)
	}

	return links, nil
}

func (r *linkRepository) Delete(ctx context.Context, id uint) store.StoreError {
	q := "UPDATE links SET is_deleted = true WHERE link_id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, id); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *linkRepository) DeletePermanently(ctx context.Context, id uint) store.StoreError {
	q := "DELETE FROM links WHERE link_id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, id); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *linkRepository) Edit(ctx context.Context, description, url string, id uint) (*models.Link, store.StoreError) {
	q := "UPDATE links SET description = $1, url = $2 WHERE link_id = $3 RETURNING link_id, url, description, created_at, modified_at"

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, description, url, id).Scan(&l.Id, &l.Url, &l.Description, &l.CreatedAt, &l.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}

func (r *linkRepository) EditUrl(ctx context.Context, url string, id uint) (*models.Link, store.StoreError) {
	q := "UPDATE links SET url = $1 WHERE link_id = $2 RETURNING link_id, url, description, created_at, modified_at"

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, url, id).Scan(&l.Id, &l.Url, &l.Description, &l.CreatedAt, &l.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}

func (r *linkRepository) EditDescription(ctx context.Context, description string, id uint) (*models.Link, store.StoreError) {
	q := "UPDATE links SET description = $1 WHERE link_id = $2 RETURNING link_id, url, description, created_at, modified_at"

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, description, id).Scan(&l.Id, &l.Url, &l.Description, &l.CreatedAt, &l.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}

func (r *linkRepository) Restore(ctx context.Context, id uint) (*models.Link, store.StoreError) {
	q := "UPDATE links SET is_deleted = false WHERE link_id = $1 RETURNING link_id, url, description, created_at, modified_at;"

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, id).Scan(&l.Id, &l.Url, &l.Description, &l.CreatedAt, &l.ModifiedAt); err != nil {
		r.logger.Errorf("Unknown Restore error: %v", err)

		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}
