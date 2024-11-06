package model

import "time"

type ReportRentModel struct {
	CustomerID  string
	FullName    string
	VehicleName string
	Price       float64
	DateIn      time.Time
	DateOut     time.Time
}
