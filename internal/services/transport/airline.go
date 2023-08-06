package transport

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int   `json:"statusCode"`
	Data       Data  `json:"data"`
	Error      error `json:"error"`
}

type Data struct {
	Airlines []*entities.Airline `json:"airlines"`
}

type UpdateAirline struct {
	Name string `json:"name"`
}

func ResponseAirlines(c *fiber.Ctx, airlines []*entities.Airline, err error) error {

	data := Data{Airlines: airlines}
	var statusCode int
	if err != nil {
		statusCode = 500
	} else {
		statusCode = 201
	}

	response := Response{
		StatusCode: statusCode,
		Data:       data,
		Error:      err,
	}
	return c.JSON(response)

}

func ResponseAirline(c *fiber.Ctx, airline *entities.Airline, err error) error {

	airlines := make([]*entities.Airline, 0)
	if airline.ID != "" {
		airlines = append(airlines, airline)
	}
	data := Data{Airlines: airlines}

	var statusCode int
	if airline.ID == "" {
		statusCode = 404
	} else if err != nil {
		statusCode = 500
	} else {
		statusCode = 201
	}

	response := Response{
		StatusCode: statusCode,
		Data:       data,
		Error:      err,
	}
	return c.JSON(response)

}

func ValidateInsertData(c *fiber.Ctx) (*entities.Airline, error) {

	airline := new(entities.Airline)
	err := c.BodyParser(airline)

	if err != nil {
		response := Response{
			StatusCode: 500,
			Data:       Data{},
			Error:      err,
		}
		return airline, c.JSON(response)
	}
	return airline, nil

}

func ValidateUpdateData(c *fiber.Ctx) (*UpdateAirline, error) {
	var updateAirlineData UpdateAirline
	err := c.BodyParser(&updateAirlineData)
	return &updateAirlineData, err
}
