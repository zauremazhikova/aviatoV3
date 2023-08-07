package handlers

import (
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transport"
	"github.com/gofiber/fiber/v2"
)

func GetAllAirlines(c *fiber.Ctx) error {

	responseAirline, err := repositories.GetAirlines()
	return transport.ResponseAirlines(c, responseAirline, err)

}

func GetSingleAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)
	return transport.ResponseAirline(c, airline, err)

}

func CreateAirline(c *fiber.Ctx) error {

	airline, err := transport.ValidateAirlineInsertData(c)
	if err != nil {
		return err
	}
	err = repositories.CreateAirline(airline)
	return transport.ResponseAirlineCreate(c, err)

}

func UpdateAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	err = transport.ResponseAirlineNotFound(c, airline, err)
	if err != nil {
		return err
	}

	updateAirlineData, err := transport.ValidateAirlineUpdateData(c)
	if err != nil {
		return err
	}

	airline.Name = updateAirlineData.Name
	err = repositories.UpdateAirline(airline)

	return transport.ResponseAirlineUpdate(c, err)
}

func DeleteAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	err = transport.ResponseAirlineNotFound(c, airline, err)
	if err != nil {
		return err
	}

	err = repositories.DeleteAirline(id)
	return transport.ResponseAirlineDelete(c, err)

}
