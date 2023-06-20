package testpsql

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestMessage_Create(t *testing.T) {
	// get user
	user := getUser(t)

	msg := models.TestMessage(t, user)

	// create message
	_, err := Store.Message.Create(bgCtx(), msg)

	assert.NoError(t, err)
}

func TestMessage_FindAll(t *testing.T) {
	// get user
	user := getUser(t)

	msg := models.TestMessage(t, user)

	// find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	var isConstains bool

	for _, _msg := range msgs {
		if _msg.Content == msg.Content {
			isConstains = true
		}
	}

	assert.True(t, isConstains)

	isConstains = false

	// find all messages
	msgs, err = Store.Message.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	for _, _msg := range msgs {
		if _msg.Content == msg.Content {
			isConstains = true
		}
	}

	assert.True(t, isConstains)
}

func TestMessage_Find(t *testing.T) {
	// get user
	user := getUser(t)

	msg := models.TestMessage(t, user)

	// find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// find test message in database
	for _, _msg := range msgs {
		if _msg.Content == msg.Content {
			msg = _msg
		}
	}

	// finc message
	_msg, err := Store.Message.Find(bgCtx(), msg.Id)

	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(_msg, msg))
}

func TestMessage_EditWithUsername(t *testing.T) {
	// get user
	user := getUser(t)

	msg := models.TestMessage(t, user)

	// find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// find test message in database
	for _, _msg := range msgs {
		if _msg.Content == msg.Content {
			msg = _msg
		}
	}

	assert.True(t, msg.Id != 0)

	msgContent := "new content with username"

	// update message content
	newMsg, err := Store.Message.EditWithUsername(bgCtx(), msgContent, msg.Id, user.Username)

	assert.NoError(t, err)

	fmt.Printf("newMsg: %v\n", newMsg)
	assert.True(t, newMsg.Content == msgContent)
}

func TestMessage_Edit(t *testing.T) {
	// get user
	user := getUser(t)

	msg := models.TestMessage(t, user)

	// find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// find test message in database
	for _, _msg := range msgs {
		if _msg.Username == msg.Username {
			msg = _msg
		}
	}

	assert.True(t, msg.Id != 0)

	msgContent := "new content"

	// update message content
	newMsg, err := Store.Message.Edit(bgCtx(), msgContent, msg.Id)

	assert.NoError(t, err)

	fmt.Printf("newMsg: %v\n", newMsg)
	assert.True(t, newMsg.Content == msgContent)
}

func TestMessage_DeleteWithUsername(t *testing.T) {
	// get user
	user := getUser(t)

	var msg *models.Message
	messageContent := "new content"

	// find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// find test message in database
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			msg = _msg
		}
	}

	assert.NotNil(t, msg)

	// delete message
	err = Store.Message.DeleteWithUsername(bgCtx(), msg.Id, user.Username)

	assert.NoError(t, err)

	// find messages
	msgs, err = Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	var isConstains bool

	// check message constains
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			isConstains = true
		}
	}

	assert.False(t, isConstains)

	// find all messages
	msgs, err = Store.Message.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	// check message constains
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			isConstains = true
		}
	}

	assert.True(t, isConstains)
}

func TestMessage_Restore(t *testing.T) {
	// get user
	// user := getUser(t)

	var msg *models.Message
	messageContent := "new content"

	// find all messages
	msgs, err := Store.Message.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	// find test message in database
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			msg = _msg
		}
	}

	assert.NotNil(t, msg)

	// restore message
	_msg, err := Store.Message.Restore(bgCtx(), msg.Id)

	assert.NoError(t, err)

	fmt.Printf("msg: %v\n", msg)
	fmt.Printf("_msg: %v\n", _msg)

	assert.True(t, _msg.Id == msg.Id && _msg.Username == msg.Username && _msg.Content == msg.Content)

	// find messages
	msgs, err = Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	var isConstains bool

	// check message constains
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			isConstains = true
		}
	}

	assert.True(t, isConstains)
}

func TestMessage_Delete(t *testing.T) {
	var msg *models.Message
	messageContent := "new content"

	// find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// find test message in database
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			msg = _msg
		}
	}

	assert.NotNil(t, msg)

	// delete message
	err = Store.Message.Delete(bgCtx(), msg.Id)

	assert.NoError(t, err)

	// find messages
	msgs, err = Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	var isConstains bool

	// check message constains
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			isConstains = true
		}
	}

	assert.False(t, isConstains)

	// find all messages
	msgs, err = Store.Message.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	// check message constains
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			isConstains = true
		}
	}

	assert.True(t, isConstains)

	// restore message
	_, err = Store.Message.Restore(bgCtx(), msg.Id)

	assert.NoError(t, err)
}

func TestMessage_DeletePermanently(t *testing.T) {
	var msg *models.Message
	messageContent := "new content"

	// find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// find test message in database
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			msg = _msg
		}
	}

	assert.NotNil(t, msg)

	// delete message
	err = Store.Message.DeletePermanently(bgCtx(), msg.Id)

	assert.NoError(t, err)

	// find all messages
	msgs, err = Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	var deletedMsg *models.Message

	// find test message in database
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			deletedMsg = _msg
		}
	}

	assert.Nil(t, deletedMsg)

	// remove user
	removeUser(t)
}
