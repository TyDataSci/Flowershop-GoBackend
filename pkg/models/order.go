package models

type Order struct {
	ID           string  `json:"id"`
	Date         string  `json:"date"`
	UserID       string  `json:"userid"`
	OrderItems   string  `json:"orderitems"`
	DeliveryType string  `json:"deliverytype"`
	Note         string  `json:"note"`
	Instructions string  `json:"instructions"`
	TotalCost    float64 `json:"totalcost"`
}
