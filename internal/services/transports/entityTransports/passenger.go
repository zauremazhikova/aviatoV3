package entityTransports

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
)

type ResponseStructurePassenger struct {
	Data    ResponseDataPassenger `json:"data"`
	Error   error                 `json:"error"`
	Message string                `json:"message"`
}

type ResponseDataPassenger struct {
	Passengers []*entities.Passenger `json:"passengers"`
}

type UpdatePassengerStructure struct {
	Name string `json:"name"`
}

type InsertPassengerStructure struct {
	Name string `json:"name"`
}

// Валидация входящих данных

func ValidatePassengerInsertData(c *fiber.Ctx) (*InsertPassengerStructure, error) {
	var insertPassengerStructure InsertPassengerStructure
	err := c.BodyParser(&insertPassengerStructure)

	if err != nil {
		return &insertPassengerStructure, c.Status(500).JSON(ResponsePassengerInputError(err))
	}
	return &insertPassengerStructure, nil
}

func ValidatePassengerUpdateData(c *fiber.Ctx) (*UpdatePassengerStructure, error) {

	var updatePassengerData UpdatePassengerStructure
	err := c.BodyParser(&updatePassengerData)

	if err != nil {
		return &updatePassengerData, c.Status(500).JSON(ResponsePassengerInputError(err))
	}
	return &updatePassengerData, nil
}

// Ответы при наличии ошибок

func ResponsePassengerNotFound(c *fiber.Ctx, passenger *entities.Passenger, err error) error {

	if err != nil || passenger.ID == "" {
		return ResponsePassenger(c, passenger, err)
	}
	return nil
}

func ResponsePassengerInputError(err error) ResponseStructurePassenger {

	response := ResponseStructurePassenger{
		Data:    ResponseDataPassenger{},
		Error:   err,
		Message: "Something wrong with your input data",
	}
	return response
}

// Ответы

func ResponsePassengers(c *fiber.Ctx, passengers []*entities.Passenger, err error) error {

	data := ResponseDataPassenger{Passengers: passengers}
	var statusCode int
	var message string
	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else if len(passengers) == 0 {
		statusCode = 404
		message = "Passengers not found"
	} else {
		statusCode = 201
		message = "Passengers found"
	}

	response := ResponseStructurePassenger{
		Data:    data,
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponsePassenger(c *fiber.Ctx, passenger *entities.Passenger, err error) error {

	passengers := make([]*entities.Passenger, 0)
	if passenger.ID != "" {
		passengers = append(passengers, passenger)
	}
	return ResponsePassengers(c, passengers, err)

}

func ResponsePassengerCreate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Passenger has created"
	}
	response := ResponseStructurePassenger{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponsePassengerUpdate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Passenger has updated"
	}
	response := ResponseStructurePassenger{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponsePassengerDelete(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Passenger has deleted"
	}
	response := ResponseStructurePassenger{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}
