package model

import "time"

type ReportCreditModel struct {
	ID       string
	FullName string
	Vehicle  string
	Price    float64
	DateIn   time.Time
	DateOut  time.Time
}
