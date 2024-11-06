package model

import "time"

type RentModel struct {
	ID         string
	VehicleID  string
	CustomerID string
	Price      float64
	DateIn     time.Time
	DateOut    time.Time
	Status     string
	CreatedBy  string
	UpdatedBy  string
	Duration   float64
}
