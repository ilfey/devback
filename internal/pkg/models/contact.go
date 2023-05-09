package models

type Contact struct {
	BaseModel
	Title string `json:"title" binding:"required"`
	Link  struct {
		BaseModel
		Description string `json:"description"`
		Url         string `json:"url"`
	} `json:"link"`
}
