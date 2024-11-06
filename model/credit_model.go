package model

import "time"

type CreditModel struct {
	ID           string
	VehicleID    string
	CustomerID   string
	Price        float64
	Interest     float64
	DateIn       time.Time
	DateOut      time.Time
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string
	Duration     float64
	CreditDetail InstallmentCreditModel
}
