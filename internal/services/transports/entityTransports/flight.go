package entityTransports

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
	"time"
)

type ResponseStructureFlight struct {
	Data    ResponseDataFlight `json:"data"`
	Error   error              `json:"error"`
	Message string             `json:"message"`
}

type ResponseDataFlight struct {
	Flights []*entities.Flight `json:"flights"`
}

type UpdateFlightStructure struct {
	FlightNumber  string    `json:"flightNumber"`
	DirectionID   string    `json:"directionID"`
	DepartureTime time.Time `json:"departureTime"`
	ArrivalTime   time.Time `json:"arrivalTime"`
	SeatsNumber   int       `json:"seatsNumber"`
	Price         float64   `json:"price"`
}

type InsertFlightStructure struct {
	FlightNumber  string    `json:"flightNumber"`
	DirectionID   string    `json:"directionID"`
	DepartureTime time.Time `json:"departureTime"`
	ArrivalTime   time.Time `json:"arrivalTime"`
	SeatsNumber   int       `json:"seatsNumber"`
	Price         float64   `json:"price"`
}

// Валидация входящих данных

func ValidateFlightInsertData(c *fiber.Ctx) (*InsertFlightStructure, error) {
	var insertFlightStructure InsertFlightStructure
	err := c.BodyParser(&insertFlightStructure)

	if err != nil {
		return &insertFlightStructure, c.Status(500).JSON(ResponseFlightInputError(err))
	}
	return &insertFlightStructure, nil
}

func ValidateFlightUpdateData(c *fiber.Ctx) (*UpdateFlightStructure, error) {

	var updateFlightData UpdateFlightStructure
	err := c.BodyParser(&updateFlightData)

	if err != nil {
		return &updateFlightData, c.Status(500).JSON(ResponseFlightInputError(err))
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
		Data:    ResponseDataFlight{},
		Error:   err,
		Message: "Something wrong with your input data",
	}
	return response
}

func ResponseFlightDirectionNotFound(c *fiber.Ctx, direction *entities.Direction, err error) error {
	if err != nil || direction.ID == "" {
		response := ResponseStructureFlight{
			Error:   err,
			Message: "Direction not found",
		}
		return c.Status(404).JSON(response)
	}
	return nil
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
		Data:    data,
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

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
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

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
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

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
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}
