package entityTransports

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
)

type ResponseStructureBooking struct {
	StatusCode int                 `json:"statusCode"`
	Data       ResponseDataBooking `json:"data"`
	Error      error               `json:"error"`
	Message    string              `json:"message"`
}

type ResponseDataBooking struct {
	Bookings []*entities.Booking `json:"bookings"`
}

type UpdateBookingStructure struct {
	BookingNumber string `json:"bookingNumber"`
	FlightID      string `json:"flightID"`
	PassengerID   string `json:"passengerID"`
}

type InsertBookingStructure struct {
	BookingNumber string `json:"bookingNumber"`
	FlightID      string `json:"flightID"`
	PassengerID   string `json:"passengerID"`
}

// Валидация входящих данных

func ValidateBookingInsertData(c *fiber.Ctx) (*InsertBookingStructure, error) {
	var insertBookingStructure InsertBookingStructure
	err := c.BodyParser(insertBookingStructure)

	if err != nil {
		return &insertBookingStructure, c.JSON(ResponseBookingInputError(err))
	}
	return &insertBookingStructure, nil
}

func ValidateBookingUpdateData(c *fiber.Ctx) (*UpdateBookingStructure, error) {

	var updateBookingData UpdateBookingStructure
	err := c.BodyParser(&updateBookingData)

	if err != nil {
		return &updateBookingData, c.JSON(ResponseBookingInputError(err))
	}
	return &updateBookingData, nil
}

// Ответы при наличии ошибок

func ResponseBookingNotFound(c *fiber.Ctx, booking *entities.Booking, err error) error {

	if err != nil || booking.ID == "" {
		return ResponseBooking(c, booking, err)
	}
	return nil
}

func ResponseBookingInputError(err error) ResponseStructureBooking {

	response := ResponseStructureBooking{
		StatusCode: 500,
		Data:       ResponseDataBooking{},
		Error:      err,
		Message:    "Something wrong with your input data",
	}
	return response
}

func ResponseBookingPassengerNotFound(c *fiber.Ctx, passenger *entities.Passenger, err error) error {
	if err != nil || passenger.ID == "" {
		response := ResponseStructureBooking{
			Error:   err,
			Message: "Passenger not found",
		}
		return c.Status(404).JSON(response)
	}
	return nil
}

func ResponseBookingFlightNotFound(c *fiber.Ctx, flight *entities.Flight, err error) error {
	if err != nil || flight.ID == "" {
		response := ResponseStructureBooking{
			Error:   err,
			Message: "Flight not found",
		}
		return c.Status(404).JSON(response)
	}
	return nil
}

// Ответы

func ResponseBookings(c *fiber.Ctx, bookings []*entities.Booking, err error) error {

	data := ResponseDataBooking{Bookings: bookings}
	var statusCode int
	var message string
	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else if len(bookings) == 0 {
		statusCode = 404
		message = "Bookings not found"
	} else {
		statusCode = 201
		message = "Bookings found"
	}

	response := ResponseStructureBooking{
		StatusCode: statusCode,
		Data:       data,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}

func ResponseBooking(c *fiber.Ctx, booking *entities.Booking, err error) error {

	bookings := make([]*entities.Booking, 0)
	if booking.ID != "" {
		bookings = append(bookings, booking)
	}
	return ResponseBookings(c, bookings, err)

}

func ResponseBookingCreate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Booking has created"
	}
	response := ResponseStructureBooking{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}

func ResponseBookingUpdate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "v has updated"
	}
	response := ResponseStructureBooking{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}

func ResponseBookingDelete(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Booking has deleted"
	}
	response := ResponseStructureBooking{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}
