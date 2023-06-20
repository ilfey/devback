package testpsql

import (
	"testing"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestUser_Create(t *testing.T) {
	u := models.TestUser(t)

	// create user
	_, err := Store.User.Create(bgCtx(), u)

	assert.NoError(t, err)
}

func TestUser_Find(t *testing.T) {
	u := models.TestUser(t)

	// find user
	user, err := Store.User.Find(bgCtx(), u.Username)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	// comparing attrs
	assert.True(t, u.Username == user.Username)
	assert.True(t, user.ComparePassword(u.Password))
}

func TestUser_ResetPassword(t *testing.T) {
	u := models.TestUser(t)
	u.Password = "Pa$$w0rd"

	// reset user password
	_, err := Store.User.ResetPassword(bgCtx(), u)
	assert.NoError(t, err)

	// find user
	f, err := Store.User.Find(bgCtx(), u.Username)
	assert.NoError(t, err)
	assert.NotNil(t, f)

	// comparing password
	assert.True(t, f.ComparePassword(u.Password))
}

func TestUser_Delete(t *testing.T) {
	u := models.TestUser(t)

	// delete user
	err := Store.User.Delete(bgCtx(), u.Username)

	assert.NoError(t, err)
}

func TestUser_Restore(t *testing.T) {
	u := models.TestUser(t)

	// restore user
	user, err := Store.User.Restore(bgCtx(), u.Username)

	assert.NoError(t, err)

	assert.True(t, user.Username == u.Username)
}

func TestUser_DeletePermanently(t *testing.T) {
	u := models.TestUser(t)

	// delete permanently user
	err := Store.User.DeletePermanently(bgCtx(), u.Username)
	assert.NoError(t, err)

	// find deleted user
	_, err = Store.User.Find(bgCtx(), u.Username)
	assert.Error(t, err)
}
