package testpsql

import (
	"testing"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestUser_Create(t *testing.T) {
	u := models.TestUser(t)

	err := Store.User.Create(bgCtx(), u)

	assert.NoError(t, err)
}

func TestUser_Find(t *testing.T) {
	u := models.TestUser(t)

	f, err := Store.User.Find(bgCtx(), u.Username)
	assert.NoError(t, err)
	assert.NotNil(t, f)

	assert.True(t, u.Username == f.Username)
	assert.True(t, f.ComparePassword(u.Password))
}

func TestUser_ResetPassword(t *testing.T) {
	u := models.TestUser(t)
	u.Password = "Pa$$w0rd"

	err := Store.User.ResetPassword(bgCtx(), u)
	assert.NoError(t, err)

	f, err := Store.User.Find(bgCtx(), u.Username)
	assert.NoError(t, err)
	assert.NotNil(t, f)

	assert.True(t, f.ComparePassword(u.Password))
}

func TestUser_Delete(t *testing.T) {
	u := models.TestUser(t)

	err := Store.User.Delete(bgCtx(), u.Username)

	assert.NoError(t, err)
}

func TestUser_Restore(t *testing.T) {
	u := models.TestUser(t)

	err := Store.User.Restore(bgCtx(), u.Username)

	assert.NoError(t, err)
}

func TestUser_DeletePermanently(t *testing.T) {
	u := models.TestUser(t)

	err := Store.User.DeletePermanently(bgCtx(), u.Username)
	assert.NoError(t, err)

	_, err = Store.User.Find(bgCtx(), u.Username)
	assert.Error(t, err)
}