package models

type Project struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Url         string `json:"url"`
}
