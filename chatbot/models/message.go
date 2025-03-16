package models

type Message struct {
	UserMessage string `json:"user_message" binding:"required"`
}
