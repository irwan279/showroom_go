package model

import "time"

type CustomerModel struct {
	ID        string
	User_id   string
	FullName  string
	NIK       string
	NoPhone   string
	Email     string
	Address   string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy *string
}
