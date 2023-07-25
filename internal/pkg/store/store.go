package store

import (
	"context"

	"github.com/ilfey/devback/internal/pkg/models"
)

type Store struct {
	Contact ContactRepository
	Image   ImageRepository
	Link    LinkRepository
	Message MessageRepository
	Project ProjectRepository
	User    UserRepository
}

type ContactRepository interface {
	Create(ctx context.Context, title string, linkId uint) (*models.Contact, StoreError)
	Find(ctx context.Context, id uint) (*models.Contact, StoreError)
	FindAll(ctx context.Context) ([]*models.Contact, StoreError)
	Delete(ctx context.Context, id uint) StoreError
	DeletePermanently(ctx context.Context, id uint) StoreError
	// Edit(context.Context, string, string, uint) error // TODO add meth
	// EditTitle(context.Context, string, uint) error // TODO add meth
	// EditLinkId(context.Context, string, uint) error // TODO add meth
}

type ImageRepository interface {

}

type LinkRepository interface {
	Create(ctx context.Context, link *models.Link) (*models.Link, StoreError)
	Find(ctx context.Context, id uint, isIncludeDeleted bool) (*models.Link, StoreError)
	FindAll(ctx context.Context, isIncludeDeleted bool) ([]*models.Link, StoreError)
	Delete(ctx context.Context, id uint) StoreError
	DeletePermanently(ctx context.Context, id uint) StoreError
	Edit(ctx context.Context, newDescription string, newUrl string, id uint) (*models.Link, StoreError)
	EditUrl(ctx context.Context, newUrl string, id uint) (*models.Link, StoreError)
	EditDescription(ctx context.Context, newDescription string, id uint) (*models.Link, StoreError)
	Restore(ctx context.Context, id uint) (*models.Link, StoreError)
}

type MessageRepository interface {
	Create(ctx context.Context, message *models.Message) (*models.Message, StoreError)
	Find(ctx context.Context, id uint, isIncludeDeleted bool) (*models.Message, StoreError)
	FindAll(ctx context.Context, isIncludeDeleted bool) ([]*models.Message, StoreError)
	FindAllWithScrolling(ctx context.Context, cursor int, limit int, isInverse bool, isIncludeDeleted bool) ([]*models.Message, StoreError)
	EditWithUsername(ctx context.Context, newContent string, id uint, username string) (*models.Message, StoreError)
	Edit(ctx context.Context, newContent string, id uint) (*models.Message, StoreError)
	DeleteWithUsername(ctx context.Context, id uint, username string) StoreError
	Delete(ctx context.Context, id uint) StoreError
	DeletePermanently(ctx context.Context, id uint) StoreError
	Restore(ctx context.Context, id uint) (*models.Message, StoreError)
}

type ProjectRepository interface {
	Create(ctx context.Context, project *models.Project) (*models.Project, StoreError)
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, StoreError)
	Find(ctx context.Context, username string) (*models.User, StoreError)
	ResetPassword(ctx context.Context, user *models.User) (*models.User, StoreError)
	Delete(ctx context.Context, username string) StoreError
	DeletePermanently(ctx context.Context, username string) StoreError
	Restore(ctx context.Context, username string) (*models.User, StoreError)
}
