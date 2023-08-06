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

	airline, err := transport.ValidateInsertData(c)
	if err != nil {
		return err
	}
	err = repositories.CreateAirline(airline)
	return transport.ResponseAirline(c, airline, err)

}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": err})
	}

	var updateAirlineData transport.UpdateAirline
	err = c.BodyParser(&updateAirlineData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	airline.Name = updateAirlineData.Name
	err = repositories.UpdateAirline(airline)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Airline has not updated", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Airline has Updated", "data": airline})
}

func DeleteAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	if airline.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": nil})
	}

	err = repositories.DeleteAirline(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete Airline", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Airline deleted"})
}
