package model

import "time"

type CashModel struct {
	ID          string
	VehicleID   string
	CustomerID  string
	Price       float64
	DatePayment time.Time
	CreatedAt   time.Time
	CreatedBy   string
}
