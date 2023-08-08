package transport

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
)

type ResponseStructureDirection struct {
	StatusCode int                   `json:"statusCode"`
	Data       ResponseDataDirection `json:"data"`
	Error      error                 `json:"error"`
	Message    string                `json:"message"`
}

type ResponseDataDirection struct {
	Directions []*entities.Direction `json:"directions"`
}

type UpdateDirection struct {
	OriginCityID      string `json:"originCityID"`
	DestinationCityID string `json:"destinationCityID"`
	AirlineID         string `json:"airlineID"`
}

// Валидация входящих данных

func ValidateDirectionInsertData(c *fiber.Ctx) (*entities.Direction, error) {
	direction := new(entities.Direction)
	err := c.BodyParser(direction)

	if err != nil {
		return direction, c.JSON(ResponseDirectionInputError(err))
	}
	return direction, nil
}

func ValidateDirectionUpdateData(c *fiber.Ctx) (*UpdateDirection, error) {

	var updateDirectionData UpdateDirection
	err := c.BodyParser(&updateDirectionData)

	if err != nil {
		return &updateDirectionData, c.JSON(ResponseDirectionInputError(err))
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
		StatusCode: 500,
		Data:       ResponseDataDirection{},
		Error:      err,
		Message:    "Something wrong with your input data",
	}
	return response
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
		StatusCode: statusCode,
		Data:       data,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

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
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

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
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

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
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}
