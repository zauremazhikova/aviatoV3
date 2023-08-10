package repository

import (
	"aviatoV3/internal/database"
	"aviatoV3/internal/entity"
	"time"
)

type AirlineRepository interface {
	GetSingle(id int) (*entity.Airline, error)
	GetAll() ([]*entity.Airline, error)
	CreateUser(a *entity.Airline) error
	DeleteUser(id int) error
}

func GetAirlines() (a []*entity.Airline, err error) {
	airlines := make([]*entity.Airline, 0)

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM airlines WHERE DELETED_AT IS NULL")

	_ = db.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var airline entity.Airline
		err := rows.Scan(&airline.ID, &airline.Name, &airline.CreatedAt, &airline.UpdatedAt, &airline.DeletedAt)
		if err != nil {
			return airlines, err
		} else {
			airlines = append(airlines, &airline)
		}
	}
	return airlines, nil
}

func GetAirline(id string) (*entity.Airline, error) {

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM airlines WHERE DELETED_AT IS NULL AND ID = $1", id)
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	var airline entity.Airline
	for rows.Next() {
		err := rows.Scan(&airline.ID, &airline.Name, &airline.CreatedAt, &airline.UpdatedAt, &airline.DeletedAt)
		if err != nil {
			return &entity.Airline{}, err
		}
	}
	return &airline, nil
}

func CreateAirline(airline *entity.Airline) error {
	db := database.DB()
	_, err := db.Query("INSERT INTO airlines (name, created_at) VALUES ($1, $2)", airline.Name, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func UpdateAirline(airline *entity.Airline) error {
	db := database.DB()
	_, err := db.Query("UPDATE airlines SET name = $2, updated_at = $3 WHERE deleted_at IS NULL AND id = $1", airline.ID, airline.Name, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func DeleteAirline(id string) error {
	db := database.DB()
	_, err := db.Query("UPDATE airlines SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}
