package models

import "time"

type Session struct {
	Token   string    `json:"token"`
	UserID  int       `json:"userid"`
	OrderID int       `json:"orderid"`
	Expiry  time.Time `json:"expiry"`
}

func (s Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
}
