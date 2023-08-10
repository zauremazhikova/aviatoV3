package entityHandlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transports/entityTransports"
	"github.com/gofiber/fiber/v2"
)

func GetAllDirections(c *fiber.Ctx) error {

	responseDirections, err := repositories.GetDirections()
	return entityTransports.ResponseDirections(c, responseDirections, err)

}

func GetSingleDirection(c *fiber.Ctx) error {

	id := c.Params("id")
	direction, err := repositories.GetDirection(id)
	return entityTransports.ResponseDirection(c, direction, err)

}

func CreateDirection(c *fiber.Ctx) error {

	insertStruct, err := entityTransports.ValidateDirectionInsertData(c)
	if err != nil {
		return err
	}

	originCity, err := repositories.GetCity(insertStruct.OriginCityID)
	err = entityTransports.ResponseDirectionCityNotFound(c, originCity, err)
	if err != nil {
		return err
	}

	destinationCity, err := repositories.GetCity(insertStruct.DestinationCityID)
	err = entityTransports.ResponseDirectionCityNotFound(c, destinationCity, err)
	if err != nil {
		return err
	}

	airline, err := repositories.GetAirline(insertStruct.AirlineID)
	err = entityTransports.ResponseDirectionAirlineNotFound(c, airline, err)
	if err != nil {
		return err
	}

	direction := new(entities.Direction)
	direction.OriginCity = *originCity
	direction.DestinationCity = *destinationCity
	direction.Airline = *airline

	err = repositories.CreateDirection(direction)
	return entityTransports.ResponseDirectionCreate(c, err)

}

func UpdateDirection(c *fiber.Ctx) error {

	id := c.Params("id")
	direction, err := repositories.GetDirection(id)

	err = entityTransports.ResponseDirectionNotFound(c, direction, err)
	if err != nil {
		return err
	}

	updateStruct, err := entityTransports.ValidateDirectionUpdateData(c)
	if err != nil {
		return err
	}

	originCity, err := repositories.GetCity(updateStruct.OriginCityID)
	err = entityTransports.ResponseDirectionCityNotFound(c, originCity, err)
	if err != nil {
		return err
	}

	destinationCity, err := repositories.GetCity(updateStruct.DestinationCityID)
	err = entityTransports.ResponseDirectionCityNotFound(c, destinationCity, err)
	if err != nil {
		return err
	}

	airline, err := repositories.GetAirline(updateStruct.AirlineID)
	err = entityTransports.ResponseDirectionAirlineNotFound(c, airline, err)
	if err != nil {
		return err
	}

	direction.OriginCity = *originCity
	direction.DestinationCity = *destinationCity
	direction.Airline = *airline
	err = repositories.UpdateDirection(direction)

	return entityTransports.ResponseDirectionUpdate(c, err)
}

func DeleteDirection(c *fiber.Ctx) error {

	id := c.Params("id")
	direction, err := repositories.GetDirection(id)

	err = entityTransports.ResponseDirectionNotFound(c, direction, err)
	if err != nil {
		return err
	}

	err = repositories.DeleteDirection(id)
	return entityTransports.ResponseDirectionDelete(c, err)

}
