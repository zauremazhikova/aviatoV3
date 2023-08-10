package repository

import (
	"aviatoV3/internal/database"
	"aviatoV3/internal/entity"
	"time"
)

func GetFlights() (a []*entity.Flight, err error) {
	flights := make([]*entity.Flight, 0)

	db := database.DB()
	rows, err := db.Query("SELECT ID, FLIGHT_NUMBER, DIRECTION_ID, DEPARTURE_TIME, ARRIVAL_TIME, SEATS_NUMBER, PRICE, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM flights")
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var flight entity.Flight
		var directionID string
		err := rows.Scan(&flight.ID, &flight.FlightNumber, &directionID, &flight.DepartureTime, &flight.ArrivalTime, &flight.SeatsNumber, &flight.Price, &flight.CreatedAt, &flight.UpdatedAt, &flight.DeletedAt)
		if err != nil {
			return flights, err
		} else {
			currDirection, _ := GetDirection(directionID)
			flight.Direction = *currDirection
			flights = append(flights, &flight)
		}
	}

	return flights, nil
}

func GetFlight(id string) (*entity.Flight, error) {

	db := database.DB()
	rows, err := db.Query("SELECT ID, FLIGHT_NUMBER, DIRECTION_ID, DEPARTURE_TIME, ARRIVAL_TIME, SEATS_NUMBER, PRICE, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM flights WHERE ID = $1", id)
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	var flight entity.Flight
	var directionID string
	for rows.Next() {
		err := rows.Scan(&flight.ID, &flight.FlightNumber, &directionID, &flight.DepartureTime, &flight.ArrivalTime, &flight.SeatsNumber, &flight.Price, &flight.CreatedAt, &flight.UpdatedAt, &flight.DeletedAt)
		if err != nil {
			return &entity.Flight{}, err
		}
	}
	currDirection, _ := GetDirection(directionID)
	flight.Direction = *currDirection

	return &flight, nil
}

func GetFlightsByOriginCity(originCityID string) (a []*entity.Flight, err error) {
	flights := make([]*entity.Flight, 0)

	db := database.DB()
	rows, err := db.Query("SELECT f.ID, f.FLIGHT_NUMBER, f.DIRECTION_ID, f.DEPARTURE_TIME, f.ARRIVAL_TIME, f.SEATS_NUMBER, f.PRICE, f.CREATED_AT, COALESCE(f.UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(f.DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM flights AS f JOIN directions AS d ON f.direction_id = d.ID WHERE d.origin_city_id = $1", originCityID)
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var flight entity.Flight
		var directionID string
		err := rows.Scan(&flight.ID, &flight.FlightNumber, &directionID, &flight.DepartureTime, &flight.ArrivalTime, &flight.SeatsNumber, &flight.Price, &flight.CreatedAt, &flight.UpdatedAt, &flight.DeletedAt)
		if err != nil {
			return flights, err
		} else {
			currDirection, _ := GetDirection(directionID)
			flight.Direction = *currDirection
			flights = append(flights, &flight)
		}
	}

	return flights, nil
}

func CreateFlight(flight *entity.Flight) error {
	db := database.DB()
	_, err := db.Query("INSERT INTO flights (FLIGHT_NUMBER, DIRECTION_ID, DEPARTURE_TIME, ARRIVAL_TIME, SEATS_NUMBER, PRICE, CREATED_AT) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		flight.FlightNumber, flight.Direction.ID, flight.DepartureTime, flight.ArrivalTime, flight.SeatsNumber, flight.Price, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func UpdateFlight(flight *entity.Flight) error {
	db := database.DB()
	_, err := db.Query("UPDATE flights SET FLIGHT_NUMBER = $2, DIRECTION_ID = $3, DEPARTURE_TIME = $4, ARRIVAL_TIME = $5, SEATS_NUMBER = $6, PRICE = $7, UPDATED_AT = $8 WHERE id = $1",
		flight.ID, flight.FlightNumber, flight.Direction.ID, flight.DepartureTime, flight.ArrivalTime, flight.SeatsNumber, flight.Price, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func DeleteFlight(id string) error {
	db := database.DB()
	_, err := db.Query("UPDATE flights SET DELETED_AT = $1 WHERE id = $2", time.Now(), id)

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}
