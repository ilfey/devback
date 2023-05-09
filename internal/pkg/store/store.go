package store

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
)

type Store struct {
	User    UserRepository
	Message MessageRepository
	Link    LinkRepository
	Contact ContactRepository
}

type UserRepository interface {
	Create(context.Context, *models.User) error
	Find(context.Context, string) (*models.User, error)
	ResetPassword(context.Context, *models.User) error
	Delete(context.Context, string) error
	DeletePermanently(context.Context, string) error
	Restore(context.Context, string) error
}

type MessageRepository interface {
	Create(context.Context, *models.Message) error
	FindAll(context.Context) ([]*models.Message, error)
	EditWithUsername(context.Context, string, uint, string) error
	Edit(context.Context, string, uint) error
	DeleteWithUsername(context.Context, uint, string) error
	Delete(context.Context, uint) error
	DeletePermanently(context.Context, uint) error
	Restore(context.Context, uint) error
}

type LinkRepository interface {
	Create(context.Context, *models.Link) error
	Find(context.Context, uint) (*models.Link, error)
	FindAll(context.Context, bool) ([]*models.Link, error)
	Delete(context.Context, uint) error
	DeletePermanently(context.Context, uint) error
	Edit(context.Context, string, string, uint) error
	EditUrl(context.Context, string, uint) error
	EditDescription(context.Context, string, uint) error
	Restore(context.Context, uint) error
}

type ContactRepository interface {
	Create(context.Context, string, uint) error
	Find(context.Context, uint) (*models.Contact, error)
	FindAll(context.Context) ([]*models.Contact, error)
	Delete(ctx context.Context, id uint) error
	DeletePermanently(ctx context.Context, id uint) error
	// Edit(context.Context, string, string, uint) error // TODO add meth
	// EditTitle(context.Context, string, uint) error // TODO add meth
	// EditLinkId(context.Context, string, uint) error // TODO add meth
}
