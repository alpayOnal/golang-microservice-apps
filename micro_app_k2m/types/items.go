package types

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Price       float32 `json:"price"`
	Currency     string `json:"currency"`
}
