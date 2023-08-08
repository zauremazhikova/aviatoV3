package handlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transport"
	"github.com/gofiber/fiber/v2"
)

func GetAllPassengers(c *fiber.Ctx) error {

	responsePassenger, err := repositories.GetPassengers()
	return transport.ResponsePassengers(c, responsePassenger, err)

}

func GetSinglePassenger(c *fiber.Ctx) error {

	id := c.Params("id")
	passenger, err := repositories.GetPassenger(id)
	return transport.ResponsePassenger(c, passenger, err)

}

func CreatePassenger(c *fiber.Ctx) error {

	insertStruct, err := transport.ValidatePassengerInsertData(c)
	if err != nil {
		return err
	}
	passenger := new(entities.Passenger)
	passenger.Name = insertStruct.Name

	err = repositories.CreatePassenger(passenger)
	return transport.ResponsePassengerCreate(c, err)

}

func UpdatePassenger(c *fiber.Ctx) error {

	id := c.Params("id")
	passenger, err := repositories.GetPassenger(id)

	err = transport.ResponsePassengerNotFound(c, passenger, err)
	if err != nil {
		return err
	}

	updatePassengerData, err := transport.ValidatePassengerUpdateData(c)
	if err != nil {
		return err
	}

	passenger.Name = updatePassengerData.Name
	err = repositories.UpdatePassenger(passenger)

	return transport.ResponsePassengerUpdate(c, err)
}

func DeletePassenger(c *fiber.Ctx) error {

	id := c.Params("id")
	passenger, err := repositories.GetPassenger(id)

	err = transport.ResponsePassengerNotFound(c, passenger, err)
	if err != nil {
		return err
	}

	err = repositories.DeletePassenger(id)
	return transport.ResponsePassengerDelete(c, err)

}
