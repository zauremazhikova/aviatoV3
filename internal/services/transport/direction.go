package transport

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
)

type ResponseStructureDirection struct {
	Data    ResponseDataDirection `json:"data"`
	Error   error                 `json:"error"`
	Message string                `json:"message"`
}

type ResponseDataDirection struct {
	Directions []*entities.Direction `json:"directions"`
}

type UpdateDirectionStructure struct {
	OriginCityID      string `json:"originCityID"`
	DestinationCityID string `json:"destinationCityID"`
	AirlineID         string `json:"airlineID"`
}

type InsertDirectionStructure struct {
	OriginCityID      string `json:"originCityID"`
	DestinationCityID string `json:"destinationCityID"`
	AirlineID         string `json:"airlineID"`
}

// Валидация входящих данных

func ValidateDirectionInsertData(c *fiber.Ctx) (*InsertDirectionStructure, error) {
	var insertDirectionStructure InsertDirectionStructure
	err := c.BodyParser(&insertDirectionStructure)

	if err != nil {
		return &insertDirectionStructure, c.Status(500).JSON(ResponseDirectionInputError(err))
	}
	return &insertDirectionStructure, nil
}

func ValidateDirectionUpdateData(c *fiber.Ctx) (*UpdateDirectionStructure, error) {

	var updateDirectionData UpdateDirectionStructure
	err := c.BodyParser(&updateDirectionData)

	if err != nil {
		return &updateDirectionData, c.Status(500).JSON(ResponseDirectionInputError(err))
	}
	return &updateDirectionData, nil
}

// Ответы при наличии ошибок

func ResponseDirectionNotFound(c *fiber.Ctx, direction *entities.Direction, err error) error {

	if err != nil || direction.ID == "" {
		return ResponseDirection(c, direction, err)
	}
	return nil
}

func ResponseDirectionInputError(err error) ResponseStructureDirection {

	response := ResponseStructureDirection{
		Data:    ResponseDataDirection{},
		Error:   err,
		Message: "Something wrong with your input data",
	}
	return response
}

func ResponseDirectionCityNotFound(c *fiber.Ctx, city *entities.City, err error) error {
	if err != nil || city.ID == "" {
		response := ResponseStructureDirection{
			Error:   err,
			Message: "City not found",
		}
		return c.Status(404).JSON(response)
	}
	return nil
}

func ResponseDirectionAirlineNotFound(c *fiber.Ctx, airline *entities.Airline, err error) error {
	if err != nil || airline.ID == "" {
		response := ResponseStructureDirection{
			Error:   err,
			Message: "Airline not found",
		}
		return c.Status(404).JSON(response)
	}
	return nil
}

// Ответы

func ResponseDirections(c *fiber.Ctx, directions []*entities.Direction, err error) error {

	data := ResponseDataDirection{Directions: directions}
	var statusCode int
	var message string
	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else if len(directions) == 0 {
		statusCode = 404
		message = "Directions not found"
	} else {
		statusCode = 201
		message = "Directions found"
	}

	response := ResponseStructureDirection{
		Data:    data,
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseDirection(c *fiber.Ctx, direction *entities.Direction, err error) error {

	directions := make([]*entities.Direction, 0)
	if direction.ID != "" {
		directions = append(directions, direction)
	}
	return ResponseDirections(c, directions, err)

}

func ResponseDirectionCreate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Direction has created"
	}
	response := ResponseStructureDirection{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseDirectionUpdate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Direction has updated"
	}
	response := ResponseStructureDirection{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseDirectionDelete(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Direction has deleted"
	}
	response := ResponseStructureDirection{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}
