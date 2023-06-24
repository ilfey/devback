package psql

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type contactRepository struct {
	db     *pgxpool.Pool
	logger *logrus.Entry
}

func (r *contactRepository) Create(ctx context.Context, title string, linkId uint) (*models.Contact, store.StoreError) {
	q := `INSERT INTO contacts(title, fk_link_id) VALUES ($1, $2) RETURNING contact_id, title, fk_link_id, created_at, modified_at;`

	c := new(models.Contact)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, title, linkId).Scan(&c.Id, &c.Title, &c.Link.Id, &c.CreatedAt, &c.ModifiedAt); err != nil {
		r.logger.Errorf("Unknown Create error: %v", err)

		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return c, nil
}

func (r *contactRepository) Find(ctx context.Context, id uint) (*models.Contact, store.StoreError) {
	q := `SELECT contact_id, title, link_id, description, url, links.modified_at, links.created_at, contacts.modified_at, contacts.created_at FROM contacts JOIN links ON fk_link_id = link_id WHERE contact_id = $1 and contacts.is_deleted = false and links.is_deleted = false;`

	c := new(models.Contact)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, id).Scan(&c.Id, &c.Title, &c.Link.Id, &c.Link.Description, &c.Link.Url, &c.Link.ModifiedAt, &c.Link.CreatedAt, &c.ModifiedAt, &c.CreatedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return c, nil
}

func (r *contactRepository) FindAll(ctx context.Context) ([]*models.Contact, store.StoreError) {
	q := `SELECT contact_id, title, link_id, description, url, links.modified_at, links.created_at, contacts.modified_at, contacts.created_at FROM contacts JOIN links ON fk_link_id = link_id WHERE contacts.is_deleted = false and links.is_deleted = false;`

	r.logger.Tracef("SQL Query: %s", q)

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	var contacts []*models.Contact

	for rows.Next() {
		c := new(models.Contact)

		err := rows.Scan(&c.Id, &c.Title, &c.Link.Id, &c.Link.Description, &c.Link.Url, &c.Link.ModifiedAt, &c.Link.CreatedAt, &c.ModifiedAt, &c.CreatedAt)
		if err != nil {
			return nil, store.NewErrorAndLog(err, r.logger)
		}

		contacts = append(contacts, c)
	}

	return contacts, nil
}

func (r *contactRepository) Delete(ctx context.Context, id uint) store.StoreError {
	q := "UPDATE contacts SET is_deleted = true WHERE contact_id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, id); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *contactRepository) DeletePermanently(ctx context.Context, id uint) store.StoreError {
	q := "DELETE FROM contacts WHERE contact_id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	if _, err := r.db.Exec(ctx, q, id); err != nil {
		return store.NewErrorAndLog(err, r.logger)
	}

	return nil
}

func (r *contactRepository) Edit() {}
