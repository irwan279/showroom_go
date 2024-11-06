package model

import "time"

type ReportCashModel struct {
	ID           string
	CustomerName string
	VehicleName  string
	Price        float64
	DatePayment  *time.Time
}
