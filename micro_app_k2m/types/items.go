package types

import (
	"time"
)

type Item struct {
	Id          string    `json:"_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Price       float32   `json:"price"`
	Currency    string    `json:"currency"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewItem() Item {
	return Item{CreatedAt: time.Now()}
}
