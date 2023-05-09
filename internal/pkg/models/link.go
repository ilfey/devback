package models

type Link struct {
	BaseModel
	Description string `json:"description"`
	Url         string `json:"url" binding:"required,url"`
}
