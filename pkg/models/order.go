package models

type Order struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	UserID    int    `json:"userid"`
	Delivery  bool   `json:"delivery"`
	Completed bool   `json:"completed"`
}
