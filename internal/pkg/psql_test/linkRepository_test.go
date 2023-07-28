package testpsql

import (
	"reflect"
	"testing"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestLink_Create(t *testing.T) {
	// Get user
	user := getUser(t)

	l := models.TestLink(t, user)

	// Create link
	_, err := Store.Link.Create(bgCtx(), l)

	assert.NoError(t, err)
}

func TestLink_FindAll(t *testing.T) {

	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Get user
	user := getUser(t)

	l := models.TestLink(t, user)

	var isContains bool

	for _, link := range links {
		if link.Username == l.Username && link.Url == l.Url && link.Description == l.Description {
			isContains = true
		}
	}

	assert.True(t, isContains)

	// Find all links
	links, err = Store.Link.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	isContains = false

	for _, link := range links {
		if link.Username == l.Username && link.Url == l.Url && link.Description == l.Description {
			isContains = true
		}
	}

	assert.True(t, isContains)
}

func TestLink_Find(t *testing.T) {
	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Get user
	user := getUser(t)

	link := models.TestLink(t, user)

	for _, _link := range links {
		if _link.Username == link.Username && _link.Url == link.Url && _link.Description == link.Description {
			link = _link
		}
	}

	// Find link
	_link, err := Store.Link.Find(bgCtx(), link.Id, false)

	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(_link, link))
}

func TestLink_EditUrl(t *testing.T) {
	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Get user
	user := getUser(t)

	link := models.TestLink(t, user)

	for _, _link := range links {
		if _link.Url == link.Url && _link.Description == link.Description {
			link = _link
		}
	}

	linkUrl := "https://google.com"

	_link, err := Store.Link.EditUrl(bgCtx(), linkUrl, link.Id)

	assert.NoError(t, err)

	assert.True(t, _link.Url == linkUrl)
}

func TestLink_EditDescription(t *testing.T) {
	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Get user
	user := getUser(t)

	link := models.TestLink(t, user)

	for _, _link := range links {
		if _link.Username == link.Username && _link.Url == "https://google.com" && _link.Description == link.Description {
			link = _link
		}
	}

	linkDescription := "The search engine"

	_link, err := Store.Link.EditDescription(bgCtx(), linkDescription, link.Id)

	assert.NoError(t, err)

	assert.True(t, _link.Description == linkDescription)
}

func TestLink_Edit(t *testing.T) {
	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Get user
	user := getUser(t)

	link := models.TestLink(t, user)

	for _, _link := range links {
		if _link.Username == link.Username && _link.Url == "https://google.com" && _link.Description == "The search engine" {
			link = _link
		}
	}

	testLink := models.TestLink(t, user)

	_link, err := Store.Link.Edit(bgCtx(), testLink.Description, testLink.Url, link.Id)

	assert.NoError(t, err)

	assert.True(t, _link.Username == testLink.Username)
	assert.True(t, _link.Description == testLink.Description)
	assert.True(t, _link.Url == testLink.Url)
}

func TestLink_Delete(t *testing.T) {
	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Get user
	user := getUser(t)

	link := models.TestLink(t, user)

	for _, _link := range links {
		if _link.Username == link.Username && _link.Url == link.Url && _link.Description == link.Description {
			link = _link
		}
	}

	// Delete link
	err = Store.Link.Delete(bgCtx(), link.Id)

	assert.NoError(t, err)

	// Find link
	_link, err := Store.Link.Find(bgCtx(), link.Id, false)

	assert.Error(t, err)

	assert.Nil(t, _link)
}

func TestLink_Restore(t *testing.T) {
	// Find links
	links, err := Store.Link.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	// Get user
	user := getUser(t)

	link := models.TestLink(t, user)

	for _, _link := range links {
		if _link.Username == link.Username && _link.Url == link.Url && _link.Description == link.Description {
			link = _link
		}
	}

	// Restore link
	_link, err := Store.Link.Restore(bgCtx(), link.Id)

	assert.NoError(t, err)

	assert.False(t, _link.IsDeleted)
}

func TestLink_DeleteWithUsername(t *testing.T) {
	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Get user
	user := getUser(t)

	link := models.TestLink(t, user)

	for _, _link := range links {
		if _link.Username == link.Username && _link.Url == link.Url && _link.Description == link.Description {
			link = _link
		}
	}

	// Delete link
	err = Store.Link.DeleteWithUsername(bgCtx(), link.Id, user.Username)

	assert.NoError(t, err)

	// Find link
	_link, err := Store.Link.Find(bgCtx(), link.Id, false)

	assert.Error(t, err)

	assert.Nil(t, _link)

	// Restore link
	_link, err = Store.Link.Restore(bgCtx(), link.Id)

	assert.NoError(t, err)

	assert.False(t, _link.IsDeleted)
}

func TestLink_DeletePermanently(t *testing.T) {
	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Get user
	user := getUser(t)

	link := models.TestLink(t, user)

	for _, _link := range links {
		if _link.Username == link.Username && _link.Url == link.Url && _link.Description == link.Description {
			link = _link
		}
	}

	// Delete link
	err = Store.Link.DeletePermanently(bgCtx(), link.Id)

	assert.NoError(t, err)

	// Find link
	_link, err := Store.Link.Find(bgCtx(), link.Id, false)

	assert.Error(t, err)

	assert.Nil(t, _link)

	removeUser(t)
}
