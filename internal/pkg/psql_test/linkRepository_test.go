package testpsql

import (
	"reflect"
	"testing"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/stretchr/testify/assert"
)

/*
Create
Find
FindAll
Delete
DeletePermanently
Edit
EditUrl
EditDescription
Restore
*/

func TestLink_Create(t *testing.T) {
	l := models.TestLink(t)

	// Create link
	_, err := Store.Link.Create(bgCtx(), l)

	assert.NoError(t, err)
}

func TestLink_FindAll(t *testing.T) {

	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	l := models.TestLink(t)

	var isContains bool

	for _, link := range links {
		if link.Url == l.Url && link.Description == l.Description {
			isContains = true
		}
	}

	assert.True(t, isContains)

	// Find all links
	links, err = Store.Link.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	isContains = false

	for _, link := range links {
		if link.Url == l.Url && link.Description == l.Description {
			isContains = true
		}
	}

	assert.True(t, isContains)
}

func TestLink_Find(t *testing.T) {
	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	link := models.TestLink(t)

	for _, _link := range links {
		if _link.Url == link.Url && _link.Description == link.Description {
			link = _link
		}
	}

	// Find link
	_link, err := Store.Link.Find(bgCtx(), link.Id, false)

	assert.NoError(t, err)
	assert.NotNil(t, reflect.DeepEqual(_link, link))
}

func TestLink_EditUrl(t *testing.T) {
	// Find links
	links, err := Store.Link.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	link := models.TestLink(t)

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

	link := models.TestLink(t)

	for _, _link := range links {
		if _link.Url == "https://google.com" && _link.Description == link.Description {
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

	link := models.TestLink(t)

	for _, _link := range links {
		if _link.Url == "https://google.com" && _link.Description == "The search engine" {
			link = _link
		}
	}

    testLink := models.TestLink(t)

	_link, err := Store.Link.Edit(bgCtx(), testLink.Description, testLink.Url, link.Id)

	assert.NoError(t, err)

	assert.True(t, _link.Description == testLink.Description)
	assert.True(t, _link.Url == testLink.Url)
}