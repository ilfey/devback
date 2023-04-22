package store

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
)

type Store struct {
	User    UserRepository
	Message MessageRepository
}

type UserRepository interface {
	Create(context.Context, *models.User) error
	Find(context.Context, string) (*models.User, error)
	ResetPassword(context.Context, *models.User) error
	Delete(context.Context, string) error
}

type MessageRepository interface {
	Create(context.Context, *models.Message) error
	ReadAll(context.Context) ([]*models.Message, error)
	Edit()
	Delete()
}
