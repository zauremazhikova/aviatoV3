package service

import (
	"aviatoV3/internal/entity"
	"aviatoV3/internal/repository"
	"github.com/gofiber/fiber/v2"
	"time"
)

// Response structure

type FlightResponseStructure struct {
	Data    FlightResponseData `json:"data"`
	Error   error              `json:"error"`
	Message string             `json:"message"`
}

type FlightResponseData struct {
	Flights []*entity.Flight `json:"Flights"`
}

// Validation structures

type FlightUpdateStructure struct {
	FlightNumber  string    `json:"flightNumber"`
	DirectionID   string    `json:"directionID"`
	DepartureTime time.Time `json:"departureTime"`
	ArrivalTime   time.Time `json:"arrivalTime"`
	SeatsNumber   int       `json:"seatsNumber"`
	Price         float64   `json:"price"`
}

type FlightInsertStructure struct {
	FlightNumber  string    `json:"flightNumber"`
	DirectionID   string    `json:"directionID"`
	DepartureTime time.Time `json:"departureTime"`
	ArrivalTime   time.Time `json:"arrivalTime"`
	SeatsNumber   int       `json:"seatsNumber"`
	Price         float64   `json:"price"`
}

// Response making

func FlightResponse(c *fiber.Ctx, flights []*entity.Flight, err error, statusCode int, message string) error {

	response := FlightResponseStructure{
		Data:    FlightResponseData{Flights: flights},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Validation

func FlightCreateValidation(c *fiber.Ctx, flights []*entity.Flight) (*FlightInsertStructure, error) {

	var insertStructure FlightInsertStructure
	err := c.BodyParser(&insertStructure)

	if err != nil {
		err = FlightResponse(c, flights, err, 500, "Something wrong with your input data")
	}
	return &insertStructure, err

}

func FlightUpdateValidation(c *fiber.Ctx, cities []*entity.Flight) (*FlightUpdateStructure, error) {

	var updateStructure FlightUpdateStructure
	err := c.BodyParser(&updateStructure)

	if err != nil {
		err = FlightResponse(c, cities, err, 500, "Something wrong with your input data")
	}
	return &updateStructure, err

}

// Methods

func GetAllFlights(c *fiber.Ctx) error {

	flights, err := repository.GetFlights()
	if err != nil {
		return FlightResponse(c, flights, err, 500, "Unexpected error")
	} else if len(flights) == 0 {
		return FlightResponse(c, flights, err, 404, "Flights not found")
	}
	return FlightResponse(c, flights, err, 201, "Flights found")
}

func GetSingleFlight(c *fiber.Ctx) error {

	id := c.Params("id")
	flight, err := repository.GetFlight(id)

	flights := make([]*entity.Flight, 0)
	if flight.ID != "" {
		flights = append(flights, flight)
	}

	if err != nil {
		return FlightResponse(c, flights, err, 500, "Unexpected error")
	} else if len(flights) == 0 {
		return FlightResponse(c, flights, err, 404, "Flight not found")
	}

	return FlightResponse(c, flights, err, 201, "Flight found")

}

func CreateFlight(c *fiber.Ctx) error {

	flights := make([]*entity.Flight, 0)
	insertStructure, err := FlightCreateValidation(c, flights)
	if err != nil {
		return err
	}

	direction, err := repository.GetDirection(insertStructure.DirectionID)
	if err != nil || direction.ID == "" {
		return FlightResponse(c, flights, err, 404, "Direction not found")
	}

	flight := new(entity.Flight)
	flight.Direction = *direction
	flight.FlightNumber = insertStructure.FlightNumber
	flight.DepartureTime = insertStructure.DepartureTime
	flight.ArrivalTime = insertStructure.ArrivalTime
	flight.SeatsNumber = insertStructure.SeatsNumber
	flight.Price = insertStructure.Price

	err = repository.CreateFlight(flight)
	if err != nil {
		return FlightResponse(c, flights, err, 500, "Unexpected error")
	}
	flights = append(flights, flight)
	return FlightResponse(c, flights, err, 201, "Flight has created")

}

func UpdateFlight(c *fiber.Ctx) error {

	flights := make([]*entity.Flight, 0)
	updateStructure, err := FlightUpdateValidation(c, flights)
	if err != nil {
		return err
	}

	id := c.Params("id")
	flight, err := repository.GetFlight(id)
	if flight.ID == "" {
		return FlightResponse(c, flights, err, 404, "Flight not found")
	}

	direction, err := repository.GetDirection(updateStructure.DirectionID)
	if err != nil || direction.ID == "" {
		return FlightResponse(c, flights, err, 404, "Direction not found")
	}

	flight.Direction = *direction
	flight.FlightNumber = updateStructure.FlightNumber
	flight.DepartureTime = updateStructure.DepartureTime
	flight.ArrivalTime = updateStructure.ArrivalTime
	flight.SeatsNumber = updateStructure.SeatsNumber
	flight.Price = updateStructure.Price

	err = repository.UpdateFlight(flight)
	if err != nil {
		return FlightResponse(c, flights, err, 500, "Unexpected error")
	}
	flights = append(flights, flight)
	return FlightResponse(c, flights, err, 201, "Flight has updated")

}

func DeleteFlight(c *fiber.Ctx) error {

	flights := make([]*entity.Flight, 0)

	id := c.Params("id")
	flight, err := repository.GetFlight(id)

	if flight.ID == "" {
		return FlightResponse(c, flights, err, 404, "Flight not found")
	}

	err = repository.DeleteFlight(id)
	if err != nil {
		return FlightResponse(c, flights, err, 500, "Unexpected error")
	}

	return FlightResponse(c, flights, err, 201, "Flight has deleted")

}
