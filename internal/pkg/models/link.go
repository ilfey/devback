package models

type Link struct {
	BaseModel
	Username    string `json:"username"`
	Description string `json:"description"`
	Url         string `json:"url" binding:"required,url"`
}
