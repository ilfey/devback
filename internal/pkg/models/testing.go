package models

import (
	"testing"
)

// Returns non-existent test user
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Username: "test_user1",
		Password: "p@ssw0rd1",
	}
}

// Returns array of non-existent test users
func TestManyUsers(t *testing.T) []*User {
	t.Helper()

	return []*User{
		TestUser(t),
		{
			Username: "test_user2",
			Password: "p@ssw0rd3",
		},
		{
			Username: "test_user3",
			Password: "p@ssw0rd3",
		},
	}
}

// Returns non-existent test message
func TestMessage(t *testing.T, u *User) *Message {
	t.Helper()

	return &Message{
		Username: u.Username,
		Content:  "message",
	}
}

// Returns array of non-existent test messages.
func TestManyMessages(t *testing.T) []*Message {
	t.Helper()
	users := TestManyUsers(t)

	return []*Message{
		TestMessage(t, users[0]),
		{
			Username: users[1].Username,
			Content:  "test_user2's message",
		},
		{
			Username: users[2].Username,
			Content:  "test_user3's message",
		},
	}
}

// Returns non-existent test link
func TestLink(t *testing.T) *Link {
	t.Helper()

	return &Link{
		Url: "https://github.com/ilfey/devback",
		Description: "This project repository",
	}
}
