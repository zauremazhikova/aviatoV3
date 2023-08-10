package repository

import (
	"aviatoV3/internal/database"
	"aviatoV3/internal/entity"
	"fmt"
	"time"
)

func GetBookings() (a []*entity.Booking, err error) {
	bookings := make([]*entity.Booking, 0)

	db := database.DB()
	rows, err := db.Query("SELECT ID, BOOKING_NUMBER, FLIGHT_ID, PASSENGER_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM bookings WHERE DELETED_AT IS NULL")
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var booking entity.Booking
		var flightID string
		var passengerID string
		err := rows.Scan(&booking.ID, &booking.BookingNumber, &flightID, &passengerID, &booking.CreatedAt, &booking.UpdatedAt, &booking.DeletedAt)
		if err != nil {
			return bookings, err
		} else {
			currentFlight, _ := GetFlight(flightID)
			booking.Flight = *currentFlight

			currentPassenger, _ := GetPassenger(passengerID)
			booking.Passenger = *currentPassenger
			bookings = append(bookings, &booking)
		}
	}
	return bookings, nil
}

func GetBookingsByFlightID(flightID string) (a []*entity.Booking, err error) {
	bookings := make([]*entity.Booking, 0)

	db := database.DB()
	rows, err := db.Query("SELECT ID, BOOKING_NUMBER, FLIGHT_ID, PASSENGER_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM bookings WHERE DELETED_AT IS NULL AND FLIGHT_ID = $1", flightID)
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var booking entity.Booking
		var flightID string
		var passengerID string
		err := rows.Scan(&booking.ID, &booking.BookingNumber, &flightID, &passengerID, &booking.CreatedAt, &booking.UpdatedAt, &booking.DeletedAt)
		if err != nil {
			fmt.Println(err)
			return bookings, err
		} else {
			currentFlight, _ := GetFlight(flightID)
			booking.Flight = *currentFlight

			currentPassenger, _ := GetPassenger(passengerID)
			booking.Passenger = *currentPassenger

			bookings = append(bookings, &booking)
		}
	}

	return bookings, nil
}

func GetBooking(id string) (*entity.Booking, error) {

	db := database.DB()
	rows, err := db.Query("SELECT ID, BOOKING_NUMBER, FLIGHT_ID, PASSENGER_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM bookings WHERE DELETED_AT IS NULL AND ID = $1", id)
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	var booking entity.Booking
	var flightID string
	var passengerID string

	for rows.Next() {
		err := rows.Scan(&booking.ID, &booking.BookingNumber, &flightID, &passengerID, &booking.CreatedAt, &booking.UpdatedAt, &booking.DeletedAt)
		if err != nil {
			fmt.Println("fuck", err)
			return &entity.Booking{}, err
		}
	}
	currentFlight, _ := GetFlight(flightID)
	booking.Flight = *currentFlight

	currentPassenger, _ := GetPassenger(passengerID)
	booking.Passenger = *currentPassenger

	return &booking, nil

}

func CreateBooking(booking *entity.Booking) error {
	db := database.DB()
	_, err := db.Query("INSERT INTO bookings (booking_number, flight_id, passenger_id, created_at) VALUES ($1, $2, $3, $4)", booking.BookingNumber, booking.Flight.ID, booking.Passenger.ID, time.Now())
	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func UpdateBooking(booking *entity.Booking) error {
	db := database.DB()
	_, err := db.Query("UPDATE bookings SET booking_number = $2, flight_id = $3, passenger_id = $4, updated_at = $3 WHERE DELETED_AT IS NULL AND id = $1",
		booking.ID, booking.BookingNumber, booking.Flight.ID, booking.Passenger.ID, time.Now())
	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func DeleteBooking(id string) error {
	db := database.DB()
	_, err := db.Query("UPDATE bookings SET deleted_at = $1 WHERE id = $2", time.Now(), id)
	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}
