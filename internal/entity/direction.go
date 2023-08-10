package entity

import "time"

type Direction struct {
	ID              string    `json:"id"`
	OriginCity      City      `json:"originCity"`
	DestinationCity City      `json:"destinationCity"`
	Airline         Airline   `json:"airline"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
