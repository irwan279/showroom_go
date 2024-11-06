package model

import "time"

type CustomerRequestModel struct {
	ID        string
	Username  string
	Password  string
	User_id   string
	Role      string
	Active    bool
	FullName  string
	NIK       string
	NoPhone   string
	Email     string
	Address   string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	Balance   int
}
