package transport

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Data       Data   `json:"data"`
	Error      error  `json:"error"`
	Message    string `json:"message"`
}

type Data struct {
	Airlines []*entities.Airline `json:"airlines"`
}

type UpdateAirline struct {
	Name string `json:"name"`
}

func ValidateAirlineInsertData(c *fiber.Ctx) (*entities.Airline, error) {
	airline := new(entities.Airline)
	err := c.BodyParser(airline)

	if err != nil {
		return airline, c.JSON(ResponseAirlineInputError(err))
	}
	return airline, nil
}

func ValidateAirlineUpdateData(c *fiber.Ctx) (*UpdateAirline, error) {
	var updateAirlineData UpdateAirline
	err := c.BodyParser(&updateAirlineData)

	if err != nil {
		return &updateAirlineData, c.JSON(ResponseAirlineInputError(err))
	}
	return &updateAirlineData, nil
}

func ResponseAirlineNotFound(c *fiber.Ctx, airline *entities.Airline, err error) error {

	if err != nil || airline.ID == "" {
		return ResponseAirline(c, airline, err)
	}
	return nil
}

func ResponseAirlineInputError(err error) Response {

	response := Response{
		StatusCode: 500,
		Data:       Data{},
		Error:      err,
		Message:    "Something wrong with your input data",
	}
	return response
}

func ResponseAirlines(c *fiber.Ctx, airlines []*entities.Airline, err error) error {

	data := Data{Airlines: airlines}
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

	response := Response{
		StatusCode: statusCode,
		Data:       data,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

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
	response := Response{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

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
	response := Response{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

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
	response := Response{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}
