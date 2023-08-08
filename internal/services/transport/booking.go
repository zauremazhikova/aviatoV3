package transport

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

type UpdateBooking struct {
	Name      string `json:"name"`
	CountryID string `json:"countryID"`
}

// Валидация входящих данных

func ValidateBookingInsertData(c *fiber.Ctx) (*entities.Booking, error) {
	booking := new(entities.Booking)
	err := c.BodyParser(booking)

	if err != nil {
		return booking, c.JSON(ResponseBookingInputError(err))
	}
	return booking, nil
}

func ValidateBookingUpdateData(c *fiber.Ctx) (*UpdateBooking, error) {

	var updateBookingData UpdateBooking
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
