package models

type Order_Item struct {
	ID      int  `json:"id"`
	OrderID int  `json:"orderid"`
	ItemID  int  `json:"itemid"`
	Removed bool `json:"removed"`
}
