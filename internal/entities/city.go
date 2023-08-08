package entities

import "time"

type City struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Country   Country   `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
