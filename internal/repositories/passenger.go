package repositories

import (
	"aviatoV3/internal/database"
	"aviatoV3/internal/entities"
	"time"
)

func GetPassengers() (a []*entities.Passenger, err error) {
	countries := make([]*entities.Passenger, 0)

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, PASSPORT, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM passengers")
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var passenger entities.Passenger
		err := rows.Scan(&passenger.ID, &passenger.Name, &passenger.Passport, &passenger.CreatedAt, &passenger.UpdatedAt, &passenger.DeletedAt)
		if err != nil {
			return countries, err
		} else {
			countries = append(countries, &passenger)
		}
	}

	return countries, nil
}

func GetPassenger(id string) (*entities.Passenger, error) {

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, PASSPORT, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM passengers WHERE ID = $1", id)
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	var passenger entities.Passenger
	for rows.Next() {
		err := rows.Scan(&passenger.ID, &passenger.Name, &passenger.Passport, &passenger.CreatedAt, &passenger.UpdatedAt, &passenger.DeletedAt)
		if err != nil {
			return &entities.Passenger{}, err
		}
	}

	return &passenger, nil
}

func CreatePassenger(passenger *entities.Passenger) error {
	db := database.DB()
	_, err := db.Query("INSERT INTO passengers (name, passport, created_at) VALUES ($1, $2, $3)", passenger.Name, passenger.Passport, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func UpdatePassenger(passenger *entities.Passenger) error {
	db := database.DB()
	_, err := db.Query("UPDATE passengers SET name = $2, passport = $3, updated_at = $4 WHERE id = $1", passenger.ID, passenger.Name, passenger.Passport, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func DeletePassenger(id string) error {
	db := database.DB()
	_, err := db.Query("UPDATE passengers SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}
