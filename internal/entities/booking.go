package entities

import "time"

type Booking struct {
	ID            string    `json:"id"`
	BookingNumber string    `json:"bookingNumber"`
	Flight        Flight    `json:"flight"`
	Passenger     Passenger `json:"passenger"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
