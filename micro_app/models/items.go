package models

import (
	"time"
)

type Item struct {
	Id          string    `json:"_id,omitempty"`
	Title       string    `json:"title,omitempty" validate:"required"`
	Description string    `json:"description,omitempty" validate:"required"`
	Company     string    `json:"company,omitempty" validate:"required"  `
	Price       float32   `json:"price,omitempty" validate:"required" `
	Currency    string    `json:"currency,omitempty" validate:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" validate:"required"`
}

func NewItem() Item {
	return Item{CreatedAt: time.Now()}
}
