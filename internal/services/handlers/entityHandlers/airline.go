package entityHandlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transports/entityTransports"
	"github.com/gofiber/fiber/v2"
)

func GetAllAirlines(c *fiber.Ctx) error {

	responseAirline, err := repositories.GetAirlines()
	return entityTransports.ResponseAirlines(c, responseAirline, err)

}

func GetSingleAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)
	return entityTransports.ResponseAirline(c, airline, err)

}

func CreateAirline(c *fiber.Ctx) error {

	insertStruct, err := entityTransports.ValidateAirlineInsertData(c)
	if err != nil {
		return err
	}
	airline := new(entities.Airline)
	airline.Name = insertStruct.Name

	err = repositories.CreateAirline(airline)
	return entityTransports.ResponseAirlineCreate(c, err)

}

func UpdateAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	err = entityTransports.ResponseAirlineNotFound(c, airline, err)
	if err != nil {
		return err
	}

	updateAirlineData, err := entityTransports.ValidateAirlineUpdateData(c)
	if err != nil {
		return err
	}

	airline.Name = updateAirlineData.Name
	err = repositories.UpdateAirline(airline)

	return entityTransports.ResponseAirlineUpdate(c, err)
}

func DeleteAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	err = entityTransports.ResponseAirlineNotFound(c, airline, err)
	if err != nil {
		return err
	}

	err = repositories.DeleteAirline(id)
	return entityTransports.ResponseAirlineDelete(c, err)

}
