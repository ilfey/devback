package models

import "time"

type BaseModel struct {
	Id         uint       `json:"id"`
	ModifiedAt *time.Time `json:"modified_at"`
	CreatedAt  *time.Time `json:"created_at"`
}
