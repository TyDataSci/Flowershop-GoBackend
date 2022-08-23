package models

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
