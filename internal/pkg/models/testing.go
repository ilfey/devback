package models

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Username: "test_user1",
		Password: "p@ssw0rd1",
	}
}

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

func TestMessage(t *testing.T, u *User) *Message {
	t.Helper()

	return &Message{
		Username: u.Username,
		Content:  "message",
	}
}
