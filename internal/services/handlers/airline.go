package handlers

import (
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transport"
	"errors"
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

	airline, err := transport.ValidateInsertData(c)
	if err != nil {
		return err
	}
	err = repositories.CreateAirline(airline)
	return transport.ResponseAirline(c, airline, err)

}

func UpdateAirline(c *fiber.Ctx) error {

	id := c.Params("id")

	airline, err := repositories.GetAirline(id)
	if err != nil {
		return transport.ResponseAirline(c, airline, err)
	}

	if airline.ID == "" {
		return transport.ResponseAirline(c, airline, errors.New("airline not found"))
	}

	updateAirlineData, err := transport.ValidateUpdateData(c)
	if err != nil {
		return transport.ResponseAirline(c, airline, err)
	}

	airline.Name = updateAirlineData.Name
	err = repositories.UpdateAirline(airline)

	if err != nil {
		return transport.ResponseAirline(c, airline, err)
	}
	return transport.ResponseAirline(c, airline, err)
}

func DeleteAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	if airline.ID == "" {
		return transport.ResponseAirline(c, airline, errors.New("airline not found"))
	}

	err = repositories.DeleteAirline(id)
	if err != nil {
		return transport.ResponseAirline(c, airline, err)
	}

	return transport.ResponseAirline(c, airline, err)
}
