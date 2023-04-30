package models

import (
	"time"
)

type Message struct {
	Id         uint       `json:"id"`
	Username   string     `json:"username"`
	Content    string     `json:"content" binding:"required,min=1,max=2000"`
	Reply      *int       `json:"reply_to"`
	ModifiedAt *time.Time `json:"modified_at"`
	CreatedAt  *time.Time `json:"created_at"`
}
