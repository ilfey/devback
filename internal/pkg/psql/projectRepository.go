package psql

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type projectRepository struct {
	generator *QueryGenerator
	db        *pgxpool.Pool
	logger    *logrus.Entry
}

func (r *projectRepository) Create(ctx context.Context, m *models.Project) (*models.Project, store.StoreError) {
	q := r.generator.Insert([]string{
		"title",
		"description",
		"fk_source_link_id",
		"fk_url_link_id",
	})

	r.logger.Tracef("SQL Query: %s", q)

	project := new(models.Project)

	if err := r.db.QueryRow(ctx, q, m.Title, m.Description, m.Source, m.Url).Scan(&project.Id, &project.Title, &project.Description, &project.Source, &project.Url, &project.IsDeleted, &project.ModifiedAt, &project.CreatedAt); err != nil {
		return nil, store.NewErrorAndLog(err, r.logger)
	}

	return project, nil
}

func (r *projectRepository) FindAll(ctx context.Context, isIncludeDeleted bool) ([]*models.Project, store.StoreError) {
	config := SelectConfig{
		Attrs: []string{
			"project_id",
			"title",
			"description",
			"fk_source_link_id",
			"fk_url_link_id",
			"is_deleted",
			"created_at",
			"modified_at",
		},
		OrderBy: []Order{
			{
				Attr: "project_id",
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

	var projects []*models.Project

	for rows.Next() {
		project := new(models.Project)

		if err := rows.Scan(&project.Id, &project.Title, &project.Description, &project.Source, &project.Url, project.IsDeleted, &project.CreatedAt, &project.ModifiedAt); err != nil {
			return nil, store.NewErrorAndLog(err, r.logger)
		}

		projects = append(projects, project)
	}

	return projects, nil
}
