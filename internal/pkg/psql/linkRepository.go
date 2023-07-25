package psql

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type linkRepository struct {
	generator *QueryGenerator
	db        *pgxpool.Pool
	logger    *logrus.Entry
}

func (r *linkRepository) Create(ctx context.Context, m *models.Link) (*models.Link, store.StoreError) {

	q := r.generator.Insert([]string{
		"description",
		"url",
	})

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	var row pgx.Row

	if m.Description == "" {
		row = r.db.QueryRow(ctx, q, nil, m.Url)
	} else {
		row = r.db.QueryRow(ctx, q, m.Description, m.Url)
	}

	if err := row.Scan(&l.Id, &l.Url, &l.Description, &l.IsDeleted, &l.CreatedAt, &l.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}

func (r *linkRepository) Find(ctx context.Context, id uint, isIncludeDeleted bool) (*models.Link, store.StoreError) {
	config := SelectConfig{
		Attrs: []string{
			"link_id",
			"description",
			"url",
			"is_deleted",
			"created_at",
			"modified_at",
		},
		Condition: "link_id = $$ and is_deleted = false",
	}

	if isIncludeDeleted {
		config.Condition = "link_id = $$"
	}

	q := r.generator.Select(config)

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, id).Scan(&l.Id, &l.Url, &l.Description, &l.IsDeleted, &l.CreatedAt, &l.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}

func (r *linkRepository) FindAll(ctx context.Context, isIncludeDeleted bool) ([]*models.Link, store.StoreError) {
	config := SelectConfig{
		Attrs: []string{
			"link_id",
			"description",
			"url",
			"is_deleted",
			"created_at",
			"modified_at",
		},
		OrderBy: []Order{
			{
				Attr: "link_id",
			},
		},
	}

	if !isIncludeDeleted {
		config.Condition = "is_deleted = false"
	}

	q := r.generator.Select(config)

	r.logger.Tracef("SQL Query: %s", q)

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	var links []*models.Link

	for rows.Next() {
		link := new(models.Link)

		if err := rows.Scan(&link.Id, &link.Description, &link.Url, &link.IsDeleted, &link.CreatedAt, &link.ModifiedAt); err != nil {
			return nil, store.NewErrorAndLog(err, r.logger)
		}

		links = append(links, link)
	}

	return links, nil
}

func (r *linkRepository) Delete(ctx context.Context, id uint) store.StoreError {

	q := r.generator.Update(
		[]string{
			"is_deleted",
		},
		"link_id = $$",
	)

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, true, id); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *linkRepository) DeletePermanently(ctx context.Context, id uint) store.StoreError {
	q := r.generator.Delete("link_id = $1")

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, id); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *linkRepository) Edit(ctx context.Context, description, url string, id uint) (*models.Link, store.StoreError) {
	q := r.generator.Update(
		[]string{
			"description",
			"url",
		},
		"link_id = $$",
	)

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, description, url, id).Scan(&l.Id, &l.Description, &l.Url, &l.IsDeleted, &l.CreatedAt, &l.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}

func (r *linkRepository) EditUrl(ctx context.Context, url string, id uint) (*models.Link, store.StoreError) {
	q := r.generator.Update(
		[]string{
			"url",
		},
		"link_id = $$",
	)

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, url, id).Scan(&l.Id, &l.Description, &l.Url, &l.IsDeleted, &l.CreatedAt, &l.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}

func (r *linkRepository) EditDescription(ctx context.Context, description string, id uint) (*models.Link, store.StoreError) {
	q := r.generator.Update(
		[]string{
			"description",
		},
		"link_id = $$",
	)

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, description, id).Scan(&l.Id, &l.Description, &l.Url, &l.IsDeleted, &l.CreatedAt, &l.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}

func (r *linkRepository) Restore(ctx context.Context, id uint) (*models.Link, store.StoreError) {
	q := r.generator.Update(
		[]string{
			"is_deleted",
		},
		"link_id = $$",
	)

	l := new(models.Link)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, false, id).Scan(&l.Id, &l.Description, &l.Url, &l.IsDeleted, &l.CreatedAt, &l.ModifiedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return l, nil
}
