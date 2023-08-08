package entities

import "time"

type Flight struct {
	ID            string    `json:"id"`
	FlightNumber  string    `json:"flightNumber"`
	Direction     Direction `json:"direction"`
	DepartureTime time.Time `json:"departureTime"`
	ArrivalTime   time.Time `json:"arrivalTime"`
	SeatsNumber   int       `json:"seatsNumber"`
	Price         float64   `json:"price"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
