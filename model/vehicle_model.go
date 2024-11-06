package model

import "time"

type VehicleModel struct {
	ID                    string
	Name                  string
	Type                  string
	Identification_number int
	Machine_number        int
	Release               time.Time
	Price                 float64
	Status                string
	Is_available          bool
	NumberPlate           string
	STNK                  string
	NoBPKB                string
	Price_rent            float64
}
