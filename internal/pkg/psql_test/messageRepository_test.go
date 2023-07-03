package testpsql

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestMessage_Create(t *testing.T) {
	// Create many users if not exist
	_ = getManyUsers(t)

	// Get test messages
	msgs := models.TestManyMessages(t)

	for _, m := range msgs {
		// Create message
		_, err := Store.Message.Create(bgCtx(), m)

		assert.NoError(t, err)
	}
}

func TestMessage_FindAll(t *testing.T) {
	// Get user
	user := getUser(t)

	msg := models.TestMessage(t, user)

	// Find messages
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

	// Find all messages
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
	// Get user
	user := getUser(t)

	msg := models.TestMessage(t, user)

	// Find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Find test message in database
	for _, _msg := range msgs {
		if _msg.Content == msg.Content {
			msg = _msg
		}
	}

	// Finc message
	_msg, err := Store.Message.Find(bgCtx(), msg.Id)

	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(_msg, msg))
}

func TestMessage_EditWithUsername(t *testing.T) {
	// Get user
	user := getUser(t)

	msg := models.TestMessage(t, user)

	// Find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Find test message in database
	for _, _msg := range msgs {
		if _msg.Content == msg.Content {
			msg = _msg
		}
	}

	assert.True(t, msg.Id != 0)

	msgContent := "new content with username"

	// Update message content
	newMsg, err := Store.Message.EditWithUsername(bgCtx(), msgContent, msg.Id, user.Username)

	assert.NoError(t, err)

	assert.True(t, newMsg.Content == msgContent)
}

func TestMessage_Edit(t *testing.T) {
	// Get user
	user := getUser(t)

	msg := models.TestMessage(t, user)

	// Find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Find test message in database
	for _, _msg := range msgs {
		if _msg.Username == msg.Username {
			msg = _msg
		}
	}

	assert.True(t, msg.Id != 0)

	msgContent := "new content"

	// Update message content
	newMsg, err := Store.Message.Edit(bgCtx(), msgContent, msg.Id)

	assert.NoError(t, err)

	fmt.Printf("newMsg: %v\n", newMsg)
	assert.True(t, newMsg.Content == msgContent)
}

func TestMessage_DeleteWithUsername(t *testing.T) {
	// Get user
	user := getUser(t)

	var msg *models.Message
	messageContent := "new content"

	// Find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Find test message in database
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			msg = _msg
		}
	}

	assert.NotNil(t, msg)

	// Delete message
	err = Store.Message.DeleteWithUsername(bgCtx(), msg.Id, user.Username)

	assert.NoError(t, err)

	// Find messages
	msgs, err = Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	var isConstains bool

	// Check message constains
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			isConstains = true
		}
	}

	assert.False(t, isConstains)

	// Find all messages
	msgs, err = Store.Message.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	// Check message constains
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			isConstains = true
		}
	}

	assert.True(t, isConstains)
}

func TestMessage_FindAllWithPagination(t *testing.T) {
	// Find messages
	msgs, err := Store.Message.FindAllWithScrolling(bgCtx(), 0, 0, false)

	assert.NoError(t, err)
	assert.Len(t, msgs, 2)

	// Find messages with deleted
	msgs, err = Store.Message.FindAllWithScrolling(bgCtx(), 0, 0, true)

	assert.NoError(t, err)
	assert.Len(t, msgs, 3)

	// Find messages with deleted
	msgs, err = Store.Message.FindAllWithScrolling(bgCtx(), 1, 1, true)

	assert.NoError(t, err)
	assert.Len(t, msgs, 1)

	// Find messages with deleted and cursor
	msgs, err = Store.Message.FindAllWithScrolling(bgCtx(), int(msgs[0].Id), 1, true)

	assert.NoError(t, err)
	assert.Len(t, msgs, 1)
}

func TestMessage_Restore(t *testing.T) {
	// Get user
	// user := getUser(t)

	var msg *models.Message
	messageContent := "new content"

	// Find all messages
	msgs, err := Store.Message.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	// Find test message in database
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			msg = _msg
		}
	}

	assert.NotNil(t, msg)

	// Restore message
	_msg, err := Store.Message.Restore(bgCtx(), msg.Id)

	assert.NoError(t, err)

	fmt.Printf("msg: %v\n", msg)
	fmt.Printf("_msg: %v\n", _msg)

	assert.True(t, _msg.Id == msg.Id && _msg.Username == msg.Username && _msg.Content == msg.Content)

	// Find messages
	msgs, err = Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	var isConstains bool

	// Check message constains
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

	// Find messages
	msgs, err := Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	// Find test message in database
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			msg = _msg
		}
	}

	assert.NotNil(t, msg)

	// Delete message
	err = Store.Message.Delete(bgCtx(), msg.Id)

	assert.NoError(t, err)

	// Find messages
	msgs, err = Store.Message.FindAll(bgCtx(), false)

	assert.NoError(t, err)

	var isConstains bool

	// Check message constains
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			isConstains = true
		}
	}

	assert.False(t, isConstains)

	// Find all messages
	msgs, err = Store.Message.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	// Check message constains
	for _, _msg := range msgs {
		if _msg.Content == messageContent {
			isConstains = true
		}
	}

	assert.True(t, isConstains)

	// Restore message
	_, err = Store.Message.Restore(bgCtx(), msg.Id)

	assert.NoError(t, err)
}

func TestMessage_DeletePermanently(t *testing.T) {
	// Find messages
	msgs, err := Store.Message.FindAll(bgCtx(), true)

	assert.NoError(t, err)

	for _, m := range msgs {
		// Delete message
		err = Store.Message.DeletePermanently(bgCtx(), m.Id)

		assert.NoError(t, err)
	}

	// Find all messages
	msgs, err = Store.Message.FindAll(bgCtx(), true)

	fmt.Printf("msgs: %v\n", msgs)

	assert.NoError(t, err)
	assert.Len(t, msgs, 0)

	// Remove users
	removeManyUsers(t)
}
