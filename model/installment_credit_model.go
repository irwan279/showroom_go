package model

import "time"

type InstallmentCreditModel struct {
	ID              string
	VehicleID       string
	CreditID        string
	Price           float64
	TotalPaymentNow float64
	DatePayment     time.Time
	DateFinish      time.Time
	DueDate         int
	Status          bool
	Suspend         bool
	Current         float64
}
