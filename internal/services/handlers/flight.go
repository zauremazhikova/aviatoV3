package handlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transport"
	"github.com/gofiber/fiber/v2"
)

func GetAllFlights(c *fiber.Ctx) error {

	responseFlights, err := repositories.GetFlights()
	return transport.ResponseFlights(c, responseFlights, err)

}

func GetSingleFlight(c *fiber.Ctx) error {

	id := c.Params("id")
	flight, err := repositories.GetFlight(id)
	return transport.ResponseFlight(c, flight, err)

}

func CreateFlight(c *fiber.Ctx) error {

	insertStruct, err := transport.ValidateFlightInsertData(c)
	if err != nil {
		return err
	}

	direction, err := repositories.GetDirection(insertStruct.DirectionID)
	err = transport.ResponseFlightDirectionNotFound(c, direction, err)
	if err != nil {
		return err
	}

	flight := new(entities.Flight)
	flight.FlightNumber = insertStruct.FlightNumber
	flight.Direction = *direction
	flight.DepartureTime = insertStruct.DepartureTime
	flight.ArrivalTime = insertStruct.ArrivalTime
	flight.SeatsNumber = insertStruct.SeatsNumber
	flight.Price = insertStruct.Price

	err = repositories.CreateFlight(flight)
	return transport.ResponseFlightCreate(c, err)

}

func UpdateFlight(c *fiber.Ctx) error {

	id := c.Params("id")
	flight, err := repositories.GetFlight(id)

	err = transport.ResponseFlightNotFound(c, flight, err)
	if err != nil {
		return err
	}

	updateStruct, err := transport.ValidateFlightUpdateData(c)
	if err != nil {
		return err
	}

	direction, err := repositories.GetDirection(updateStruct.DirectionID)
	err = transport.ResponseFlightDirectionNotFound(c, direction, err)
	if err != nil {
		return err
	}

	flight.FlightNumber = updateStruct.FlightNumber
	flight.Direction = *direction
	flight.DepartureTime = updateStruct.DepartureTime
	flight.ArrivalTime = updateStruct.ArrivalTime
	flight.SeatsNumber = updateStruct.SeatsNumber
	flight.Price = updateStruct.Price

	err = repositories.UpdateFlight(flight)

	return transport.ResponseFlightUpdate(c, err)
}

func DeleteFlight(c *fiber.Ctx) error {

	id := c.Params("id")
	direction, err := repositories.GetDirection(id)

	err = transport.ResponseDirectionNotFound(c, direction, err)
	if err != nil {
		return err
	}

	err = repositories.DeleteFlight(id)
	return transport.ResponseFlightDelete(c, err)

}
