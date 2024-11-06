package model

import "time"

type MemberModel struct {
	ID         string
	CustomerID string
	Type       string
	Expire     time.Time
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
}
