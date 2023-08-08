package transport

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
	"time"
)

type ResponseStructureFlight struct {
	StatusCode int                `json:"statusCode"`
	Data       ResponseDataFlight `json:"data"`
	Error      error              `json:"error"`
	Message    string             `json:"message"`
}

type ResponseDataFlight struct {
	Flights []*entities.Flight `json:"flights"`
}

type UpdateFlight struct {
	FlightNumber  string    `json:"flightNumber"`
	DirectionID   string    `json:"directionID"`
	DepartureTime time.Time `json:"departureTime"`
	ArrivalTime   time.Time `json:"arrivalTime"`
	SeatsNumber   int       `json:"seatsNumber"`
	Price         float64   `json:"price"`
}

// Валидация входящих данных

func ValidateFlightInsertData(c *fiber.Ctx) (*entities.Flight, error) {
	flight := new(entities.Flight)
	err := c.BodyParser(flight)

	if err != nil {
		return flight, c.JSON(ResponseFlightInputError(err))
	}
	return flight, nil
}

func ValidateFlightUpdateData(c *fiber.Ctx) (*UpdateFlight, error) {

	var updateFlightData UpdateFlight
	err := c.BodyParser(&updateFlightData)

	if err != nil {
		return &updateFlightData, c.JSON(ResponseFlightInputError(err))
	}
	return &updateFlightData, nil
}

// Ответы при наличии ошибок

func ResponseFlightNotFound(c *fiber.Ctx, flight *entities.Flight, err error) error {

	if err != nil || flight.ID == "" {
		return ResponseFlight(c, flight, err)
	}
	return nil
}

func ResponseFlightInputError(err error) ResponseStructureFlight {

	response := ResponseStructureFlight{
		StatusCode: 500,
		Data:       ResponseDataFlight{},
		Error:      err,
		Message:    "Something wrong with your input data",
	}
	return response
}

// Ответы

func ResponseFlights(c *fiber.Ctx, flights []*entities.Flight, err error) error {

	data := ResponseDataFlight{Flights: flights}
	var statusCode int
	var message string
	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else if len(flights) == 0 {
		statusCode = 404
		message = "Flights not found"
	} else {
		statusCode = 201
		message = "Flights found"
	}

	response := ResponseStructureFlight{
		StatusCode: statusCode,
		Data:       data,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}

func ResponseFlight(c *fiber.Ctx, flight *entities.Flight, err error) error {

	flights := make([]*entities.Flight, 0)
	if flight.ID != "" {
		flights = append(flights, flight)
	}
	return ResponseFlights(c, flights, err)

}

func ResponseFlightCreate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Flight has created"
	}
	response := ResponseStructureFlight{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}

func ResponseFlightUpdate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Flight has updated"
	}
	response := ResponseStructureFlight{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}

func ResponseFlightDelete(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Flight has deleted"
	}
	response := ResponseStructureFlight{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}
