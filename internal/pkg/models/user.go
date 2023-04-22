package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username  string    `json:"username" binding:"required,alpha,min=3,max=16"`
	Password  string    `json:"password" binding:"required,alpha"`
	Hash      string    `json:"-"`
	CreatedAt time.Time `json:"created_at" timeformat:"2006-01-02"`
}

func (u *User) BeforeCreate() error {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.Hash = string(b)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password)) == nil
}
