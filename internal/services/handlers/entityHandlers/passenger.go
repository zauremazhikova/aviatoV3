package entityHandlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transports/entityTransports"
	"github.com/gofiber/fiber/v2"
)

func GetAllPassengers(c *fiber.Ctx) error {

	responsePassenger, err := repositories.GetPassengers()
	return entityTransports.ResponsePassengers(c, responsePassenger, err)

}

func GetSinglePassenger(c *fiber.Ctx) error {

	id := c.Params("id")
	passenger, err := repositories.GetPassenger(id)
	return entityTransports.ResponsePassenger(c, passenger, err)

}

func CreatePassenger(c *fiber.Ctx) error {

	insertStruct, err := entityTransports.ValidatePassengerInsertData(c)
	if err != nil {
		return err
	}
	passenger := new(entities.Passenger)
	passenger.Name = insertStruct.Name

	err = repositories.CreatePassenger(passenger)
	return entityTransports.ResponsePassengerCreate(c, err)

}

func UpdatePassenger(c *fiber.Ctx) error {

	id := c.Params("id")
	passenger, err := repositories.GetPassenger(id)

	err = entityTransports.ResponsePassengerNotFound(c, passenger, err)
	if err != nil {
		return err
	}

	updatePassengerData, err := entityTransports.ValidatePassengerUpdateData(c)
	if err != nil {
		return err
	}

	passenger.Name = updatePassengerData.Name
	err = repositories.UpdatePassenger(passenger)

	return entityTransports.ResponsePassengerUpdate(c, err)
}

func DeletePassenger(c *fiber.Ctx) error {

	id := c.Params("id")
	passenger, err := repositories.GetPassenger(id)

	err = entityTransports.ResponsePassengerNotFound(c, passenger, err)
	if err != nil {
		return err
	}

	err = repositories.DeletePassenger(id)
	return entityTransports.ResponsePassengerDelete(c, err)

}
