package models

type Item struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Image       string `json:"image"`
}
