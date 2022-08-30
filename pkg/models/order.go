package models

type Order struct {
	ID           string  `json:"id"`
	Date         string  `json:"date"`
	UserID       string  `json:"userid"`
	Items        []*Item `json:"items"`
	DeliveryType string  `json:"deliverytype"`
	Note         string  `json:"note"`
	Instructions string  `json:"instructions"`
	TotalCost    string  `json:"totalcost"`
}
