package models

type Item struct {
	ID          string     `json:"id"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
	Price       string     `json:"price"`
	Image       string     `json:"image"`
	Inventory   *Inventory `json:"inventory"`
}
