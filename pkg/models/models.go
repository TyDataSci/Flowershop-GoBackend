package models

type Item struct {
	ID          string     `json:"id"`
	ItemType    string     `json:"itemtype"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Inventory   *Inventory `json:"inventory"`
}

type Inventory struct {
	Count        int     `json:"count"`
	MinimumCount int     `json:"minimumcount"`
	Supplier     string  `json:"supplier"`
	LeadDays     int     `json:"leaddays"`
	Cost         float64 `json:"cost"`
}

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

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Account struct {
	ID            string   `json:"id"`
	User          *User    `json:"user"`
	FirstName     string   `json:"firstname"`
	LastName      string   `json:"lastname"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	PaymentMethod string   `json:"paymentmethod"`
	Orders        []*Order `json:"orders"`
}
