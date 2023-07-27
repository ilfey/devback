package models

type Link struct {
	BaseModel
	Username    string
	Description string `json:"description"`
	Url         string `json:"url" binding:"required,url"`
}
