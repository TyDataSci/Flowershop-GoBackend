package models

type Inventory struct {
	Count        int     `json:"count"`
	MinimumCount int     `json:"minimumcount"`
	Supplier     string  `json:"supplier"`
	LeadDays     int     `json:"leaddays"`
	Cost         float64 `json:"cost"`
}
