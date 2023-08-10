package entityTransports

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
)

type ResponseStructureAirline struct {
	Data    ResponseDataAirline `json:"data"`
	Error   error               `json:"error"`
	Message string              `json:"message"`
}

type ResponseDataAirline struct {
	Airlines []*entities.Airline `json:"airlines"`
}

type UpdateAirlineStructure struct {
	Name string `json:"name"`
}

type InsertAirlineStructure struct {
	Name string `json:"name"`
}

// Валидация входящих данных

func ValidateAirlineInsertData(c *fiber.Ctx) (*InsertAirlineStructure, error) {

	var insertAirlineStructure InsertAirlineStructure
	err := c.BodyParser(&insertAirlineStructure)

	if err != nil {
		return &insertAirlineStructure, c.Status(500).JSON(ResponseAirlineInputError(err))
	}
	return &insertAirlineStructure, nil
}

func ValidateAirlineUpdateData(c *fiber.Ctx) (*UpdateAirlineStructure, error) {

	var updateAirlineData UpdateAirlineStructure
	err := c.BodyParser(&updateAirlineData)

	if err != nil {
		return &updateAirlineData, c.Status(500).JSON(ResponseAirlineInputError(err))
	}
	return &updateAirlineData, nil
}

// Ответы при наличии ошибок

func ResponseAirlineNotFound(c *fiber.Ctx, airline *entities.Airline, err error) error {

	if err != nil || airline.ID == "" {
		return ResponseAirline(c, airline, err)
	}
	return nil
}

func ResponseAirlineInputError(err error) ResponseStructureAirline {

	response := ResponseStructureAirline{
		Data:    ResponseDataAirline{},
		Error:   err,
		Message: "Something wrong with your input data",
	}
	return response
}

// Ответы

func ResponseAirlines(c *fiber.Ctx, airlines []*entities.Airline, err error) error {

	data := ResponseDataAirline{Airlines: airlines}
	var statusCode int
	var message string
	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else if len(airlines) == 0 {
		statusCode = 404
		message = "Airlines not found"
	} else {
		statusCode = 201
		message = "Airlines found"
	}

	response := ResponseStructureAirline{
		Data:    data,
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseAirline(c *fiber.Ctx, airline *entities.Airline, err error) error {

	airlines := make([]*entities.Airline, 0)
	if airline.ID != "" {
		airlines = append(airlines, airline)
	}
	return ResponseAirlines(c, airlines, err)

}

func ResponseAirlineCreate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Airline has created"
	}
	response := ResponseStructureAirline{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseAirlineUpdate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Airline has updated"
	}
	response := ResponseStructureAirline{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseAirlineDelete(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Airline has deleted"
	}
	response := ResponseStructureAirline{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}
