package models

import (
	"time"
)

type Item struct {
	Id          string    `json:"_id,omitempty" `
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Company     string    `json:"company,omitempty" `
	Price       float32   `json:"price,omitempty" `
	Currency    string    `json:"currency,omitempty" `
	CreatedAt   time.Time `json:"created_at,omitempty" `
}

func NewItem() Item {
	return Item{CreatedAt: time.Now()}
}
