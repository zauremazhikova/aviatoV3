package service

import (
	"aviatoV3/internal/entity"
	"aviatoV3/internal/repository"
	"errors"
	"github.com/gofiber/fiber/v2"
)

// Response structure

type BookingResponseStructure struct {
	Data    BookingResponseData `json:"data"`
	Error   error               `json:"error"`
	Message string              `json:"message"`
}

type BookingResponseData struct {
	Bookings []*entity.Booking `json:"bookings"`
}

// Validation structures

type BookingUpdateStructure struct {
	BookingNumber string `json:"bookingNumber"`
	FlightID      string `json:"flightID"`
	PassengerID   string `json:"passengerID"`
}

type BookingInsertStructure struct {
	BookingNumber string `json:"bookingNumber"`
	FlightID      string `json:"flightID"`
	PassengerID   string `json:"passengerID"`
}

// Response making

func BookingResponse(c *fiber.Ctx, bookings []*entity.Booking, err error, statusCode int, message string) error {

	response := BookingResponseStructure{
		Data:    BookingResponseData{Bookings: bookings},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Validation

func BookingCreateValidation(c *fiber.Ctx, bookings []*entity.Booking) (*BookingInsertStructure, error) {

	var insertStructure BookingInsertStructure
	err := c.BodyParser(&insertStructure)

	if err != nil {
		err = BookingResponse(c, bookings, err, 500, "Something wrong with your input data")
	}
	return &insertStructure, err

}

func BookingUpdateValidation(c *fiber.Ctx, bookings []*entity.Booking) (*BookingUpdateStructure, error) {

	var updateStructure BookingUpdateStructure
	err := c.BodyParser(&updateStructure)

	if err != nil {
		err = BookingResponse(c, bookings, err, 500, "Something wrong with your input data")
	}
	return &updateStructure, err

}

// Checking if the flight doesn't have seats

func CheckFlightBookingAvailability(flight *entity.Flight) (bool, error) {

	currentBookings, err := repository.GetBookingsByFlightID(flight.ID)

	if err != nil {
		return false, err
	}

	if flight.SeatsNumber <= len(currentBookings) {
		return false, errors.New("flight is full")
	}

	return true, nil
}

// Methods

func GetAllBookings(c *fiber.Ctx) error {

	bookings, err := repository.GetBookings()
	if err != nil {
		return BookingResponse(c, bookings, err, 500, "Unexpected error")
	} else if len(bookings) == 0 {
		return BookingResponse(c, bookings, err, 404, "Bookings not found")
	}
	return BookingResponse(c, bookings, err, 201, "Bookings found")
}

func GetSingleBooking(c *fiber.Ctx) error {

	id := c.Params("id")
	booking, err := repository.GetBooking(id)

	bookings := make([]*entity.Booking, 0)
	if booking.ID != "" {
		bookings = append(bookings, booking)
	}

	if err != nil {
		return BookingResponse(c, bookings, err, 500, "Unexpected error")
	} else if len(bookings) == 0 {
		return BookingResponse(c, bookings, err, 404, "Booking not found")
	}

	return BookingResponse(c, bookings, err, 201, "Booking found")

}

func CreateBooking(c *fiber.Ctx) error {

	bookings := make([]*entity.Booking, 0)
	insertStructure, err := BookingCreateValidation(c, bookings)
	if err != nil {
		return err
	}

	flight, err := repository.GetFlight(insertStructure.FlightID)
	if err != nil || flight.ID == "" {
		return BookingResponse(c, bookings, err, 404, "Flight not found")
	}

	passenger, err := repository.GetPassenger(insertStructure.PassengerID)
	if err != nil || passenger.ID == "" {
		return BookingResponse(c, bookings, err, 404, "Passenger not found")
	}

	checkingAvailability, err := CheckFlightBookingAvailability(flight)
	if checkingAvailability == false {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flight is not available", "data": err})
	}

	booking := new(entity.Booking)
	booking.Flight = *flight
	booking.Passenger = *passenger

	err = repository.CreateBooking(booking)
	if err != nil {
		return BookingResponse(c, bookings, err, 500, "Unexpected error")
	}
	bookings = append(bookings, booking)
	return BookingResponse(c, bookings, err, 201, "Booking has created")

}

func UpdateBooking(c *fiber.Ctx) error {

	bookings := make([]*entity.Booking, 0)
	updateStructure, err := BookingUpdateValidation(c, bookings)
	if err != nil {
		return err
	}

	id := c.Params("id")
	booking, err := repository.GetBooking(id)
	if booking.ID == "" {
		return BookingResponse(c, bookings, err, 404, "Booking not found")
	}

	flight, err := repository.GetFlight(updateStructure.FlightID)
	if err != nil || flight.ID == "" {
		return BookingResponse(c, bookings, err, 404, "Flight not found")
	}

	passenger, err := repository.GetPassenger(updateStructure.PassengerID)
	if err != nil || passenger.ID == "" {
		return BookingResponse(c, bookings, err, 404, "Passenger not found")
	}

	booking.Flight = *flight
	booking.Passenger = *passenger

	err = repository.UpdateBooking(booking)
	if err != nil {
		return BookingResponse(c, bookings, err, 500, "Unexpected error")
	}
	bookings = append(bookings, booking)
	return BookingResponse(c, bookings, err, 201, "Booking has updated")

}

func DeleteBooking(c *fiber.Ctx) error {

	bookings := make([]*entity.Booking, 0)

	id := c.Params("id")
	booking, err := repository.GetBooking(id)

	if booking.ID == "" {
		return BookingResponse(c, bookings, err, 404, "Booking not found")
	}

	err = repository.DeleteBooking(id)
	if err != nil {
		return BookingResponse(c, bookings, err, 500, "Unexpected error")
	}

	return BookingResponse(c, bookings, err, 201, "Booking has deleted")

}
