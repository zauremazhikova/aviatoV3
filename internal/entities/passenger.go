package entities

import "time"

type Passenger struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Passport  string    `json:"passport"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
