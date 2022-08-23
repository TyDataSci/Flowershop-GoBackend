package models

type Item struct {
	ID          string     `json:"id"`
	ItemType    string     `json:"itemtype"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Inventory   *Inventory `json:"inventory"`
}
