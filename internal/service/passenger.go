package service

import (
	"aviatoV3/internal/entity"
	"aviatoV3/internal/repository"
	"github.com/gofiber/fiber/v2"
)

// Response structure

type PassengerResponseStructure struct {
	Data    PassengerResponseData `json:"data"`
	Error   error                 `json:"error"`
	Message string                `json:"message"`
}

type PassengerResponseData struct {
	Passengers []*entity.Passenger `json:"passengers"`
}

// Validation structures

type PassengerUpdateStructure struct {
	Name     string `json:"name"`
	Passport string `json:"passport"`
}

type PassengerInsertStructure struct {
	Name     string `json:"name"`
	Passport string `json:"passport"`
}

// Response making

func PassengerResponse(c *fiber.Ctx, passengers []*entity.Passenger, err error, statusCode int, message string) error {

	response := PassengerResponseStructure{
		Data:    PassengerResponseData{Passengers: passengers},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Validation

func PassengerCreateValidation(c *fiber.Ctx, passengers []*entity.Passenger) (*PassengerInsertStructure, error) {

	var insertStructure PassengerInsertStructure
	err := c.BodyParser(&insertStructure)

	if err != nil {
		err = PassengerResponse(c, passengers, err, 500, "Something wrong with your input data")
	}
	return &insertStructure, err

}

func PassengerUpdateValidation(c *fiber.Ctx, passengers []*entity.Passenger) (*PassengerUpdateStructure, error) {

	var updateStructure PassengerUpdateStructure
	err := c.BodyParser(&updateStructure)

	if err != nil {
		err = PassengerResponse(c, passengers, err, 500, "Something wrong with your input data")
	}
	return &updateStructure, err

}

// Methods

func GetAllPassengers(c *fiber.Ctx) error {

	passengers, err := repository.GetPassengers()
	if err != nil {
		return PassengerResponse(c, passengers, err, 500, "Unexpected error")
	} else if len(passengers) == 0 {
		return PassengerResponse(c, passengers, err, 404, "Passengers not found")
	}
	return PassengerResponse(c, passengers, err, 201, "Passengers found")
}

func GetSinglePassenger(c *fiber.Ctx) error {

	id := c.Params("id")
	passenger, err := repository.GetPassenger(id)

	passengers := make([]*entity.Passenger, 0)
	if passenger.ID != "" {
		passengers = append(passengers, passenger)
	}

	if err != nil {
		return PassengerResponse(c, passengers, err, 500, "Unexpected error")
	} else if len(passengers) == 0 {
		return PassengerResponse(c, passengers, err, 404, "Passenger not found")
	}

	return PassengerResponse(c, passengers, err, 201, "Passenger found")

}

func CreatePassenger(c *fiber.Ctx) error {

	passengers := make([]*entity.Passenger, 0)
	insertStructure, err := PassengerCreateValidation(c, passengers)
	if err != nil {
		return err
	}

	passenger := new(entity.Passenger)
	passenger.Name = insertStructure.Name
	passenger.Passport = insertStructure.Passport

	err = repository.CreatePassenger(passenger)
	if err != nil {
		return PassengerResponse(c, passengers, err, 500, "Unexpected error")
	}
	passengers = append(passengers, passenger)
	return PassengerResponse(c, passengers, err, 201, "Passenger has created")

}

func UpdatePassenger(c *fiber.Ctx) error {

	passengers := make([]*entity.Passenger, 0)
	updateStructure, err := PassengerUpdateValidation(c, passengers)
	if err != nil {
		return err
	}

	id := c.Params("id")
	passenger, err := repository.GetPassenger(id)
	if passenger.ID == "" {
		return PassengerResponse(c, passengers, err, 404, "Passenger not found")
	}

	passenger.Name = updateStructure.Name
	passenger.Passport = updateStructure.Passport

	err = repository.UpdatePassenger(passenger)
	if err != nil {
		return PassengerResponse(c, passengers, err, 500, "Unexpected error")
	}
	passengers = append(passengers, passenger)
	return PassengerResponse(c, passengers, err, 201, "Passenger has updated")

}

func DeletePassenger(c *fiber.Ctx) error {

	passengers := make([]*entity.Passenger, 0)

	id := c.Params("id")
	passenger, err := repository.GetPassenger(id)

	if passenger.ID == "" {
		return PassengerResponse(c, passengers, err, 404, "Passenger not found")
	}

	err = repository.DeletePassenger(id)
	if err != nil {
		return PassengerResponse(c, passengers, err, 500, "Unexpected error")
	}

	return PassengerResponse(c, passengers, err, 201, "Passenger has deleted")

}
