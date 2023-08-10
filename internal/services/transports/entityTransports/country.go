package entityTransports

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
)

type ResponseStructureCountry struct {
	Data    ResponseDataCountry `json:"data"`
	Error   error               `json:"error"`
	Message string              `json:"message"`
}

type ResponseDataCountry struct {
	Countries []*entities.Country `json:"countries"`
}

type UpdateCountryStructure struct {
	Name string `json:"name"`
}

type InsertCountryStructure struct {
	Name string `json:"name"`
}

// Валидация входящих данных

func ValidateCountryInsertData(c *fiber.Ctx) (*InsertCountryStructure, error) {
	var insertCountryStructure InsertCountryStructure
	err := c.BodyParser(&insertCountryStructure)

	if err != nil {
		return &insertCountryStructure, c.Status(500).JSON(ResponseCountryInputError(err))
	}
	return &insertCountryStructure, nil
}

func ValidateCountryUpdateData(c *fiber.Ctx) (*UpdateCountryStructure, error) {

	var updateCountryData UpdateCountryStructure
	err := c.BodyParser(&updateCountryData)

	if err != nil {
		return &updateCountryData, c.Status(500).JSON(ResponseCountryInputError(err))
	}
	return &updateCountryData, nil
}

// Ответы при наличии ошибок

func ResponseCountryNotFound(c *fiber.Ctx, country *entities.Country, err error) error {

	if err != nil || country.ID == "" {
		return ResponseCountry(c, country, err)
	}
	return nil
}

func ResponseCountryInputError(err error) ResponseStructureCountry {

	response := ResponseStructureCountry{
		Data:    ResponseDataCountry{},
		Error:   err,
		Message: "Something wrong with your input data",
	}
	return response
}

// Ответы

func ResponseCountries(c *fiber.Ctx, countries []*entities.Country, err error) error {

	data := ResponseDataCountry{Countries: countries}
	var statusCode int
	var message string
	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else if len(countries) == 0 {
		statusCode = 404
		message = "Countries not found"
	} else {
		statusCode = 201
		message = "Countries found"
	}

	response := ResponseStructureCountry{
		Data:    data,
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseCountry(c *fiber.Ctx, country *entities.Country, err error) error {

	countries := make([]*entities.Country, 0)
	if country.ID != "" {
		countries = append(countries, country)
	}
	return ResponseCountries(c, countries, err)

}

func ResponseCountryCreate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Country has created"
	}
	response := ResponseStructureCountry{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseCountryUpdate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Country has updated"
	}
	response := ResponseStructureCountry{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseCountryDelete(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "Country has deleted"
	}
	response := ResponseStructureCountry{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}
