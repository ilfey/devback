package models

type Message struct {
	BaseModel
	Username string `json:"username"`
	Content  string `json:"content" binding:"required,min=1,max=2000"`
	Reply    *uint  `json:"reply_to"`
}
