package psql

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type contactRepository struct {
	db     *pgx.Conn
	logger *logrus.Entry
}

func (r *contactRepository) Find(ctx context.Context, id uint) (*models.Contact, error) {
	q := `SELECT c.id, c.title, l.id, l.description, l.url, l.modified_at, l.created_at, c.modified_at, c.created_at FROM contacts c JOIN links l ON c.linkId = l.id WHERE c.id = $1 and c.is_deleted = false and l.is_deleted = false;`

	c := new(models.Contact)

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, id).Scan(&c.Id, &c.Title, &c.Link.Id, &c.Link.Description, &c.Link.Url, &c.Link.ModifiedAt, &c.Link.CreatedAt, &c.ModifiedAt, &c.CreatedAt); err != nil {
		r.logger.Errorf("Unknown Find error: %v", err)

		return nil, err
	}

	return c, nil
}

func (r *contactRepository) FindAll(ctx context.Context) ([]*models.Contact, error) {
	q := `SELECT c.id, c.title, l.id, l.description, l.url, l.modified_at, l.created_at, c.modified_at, c.created_at FROM contacts c JOIN links l ON c.linkId = l.id and c.is_deleted = false and l.is_deleted = false;`

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

	var contacts []*models.Contact

	for rows.Next() {
		c := new(models.Contact)

		err := rows.Scan(&c.Id, &c.Title, &c.Link.Id, &c.Link.Description, &c.Link.Url, &c.Link.ModifiedAt, &c.Link.CreatedAt, &c.ModifiedAt, &c.CreatedAt)
		if err != nil {
			r.logger.Errorf("Scan ReadAll error: %v", err)

			return nil, err
		}

		contacts = append(contacts, c)
	}

	return contacts, nil
}

func (r *contactRepository) Create(ctx context.Context, title string, linkId uint) error {
	q := `INSERT INTO contacts(title, linkId) VALUES ($1, $2);`

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, title, linkId)
	if err != nil {
		r.logger.Errorf("Unknown Create error: %v", err)

		return err
	}

	return nil
}

func (r *contactRepository) Delete(ctx context.Context, id uint) error {
	q := "UPDATE contacts SET is_deleted = true WHERE id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		r.logger.Errorf("Unknown Delete error: %v", err)

		return err
	}

	return nil
}

func (r *contactRepository) DeletePermanently(ctx context.Context, id uint) error {
	q := "DELETE FROM contacts WHERE id = $1;"

	r.logger.Tracef("SQL Query: %s", q)

	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		r.logger.Errorf("Unknown DeletePermanently error: %v", err)

		return err
	}

	return nil
}

func (r *contactRepository) Edit() {

}
