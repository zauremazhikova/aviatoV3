package entityTransports

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
)

type ResponseCityStructure struct {
	Data    ResponseDataCity `json:"data"`
	Error   error            `json:"error"`
	Message string           `json:"message"`
}

type ResponseDataCity struct {
	Cities []*entities.City `json:"cities"`
}

type UpdateCityStructure struct {
	Name      string `json:"name"`
	CountryID string `json:"countryID"`
}

type InsertCityStructure struct {
	Name      string `json:"name"`
	CountryID string `json:"countryID"`
}

// Валидация входящих данных

func ValidateCityInsertData(c *fiber.Ctx) (*InsertCityStructure, error) {
	var insertCityData InsertCityStructure
	err := c.BodyParser(&insertCityData)
	if err != nil {
		return &insertCityData, c.JSON(ResponseCityInputError(c, err))
	}
	return &insertCityData, nil
}

func ValidateCityUpdateData(c *fiber.Ctx) (*UpdateCityStructure, error) {

	var updateCityData UpdateCityStructure
	err := c.BodyParser(&updateCityData)

	if err != nil {
		return &updateCityData, c.JSON(ResponseCityInputError(c, err))
	}
	return &updateCityData, nil
}

// Ответы при наличии ошибок

func ResponseCityNotFound(c *fiber.Ctx, city *entities.City, err error) error {

	if err != nil || city.ID == "" {
		return ResponseCity(c, city, err)
	}
	return nil
}

func ResponseCityInputError(c *fiber.Ctx, err error) error {

	response := ResponseCityStructure{
		Data:    ResponseDataCity{},
		Error:   err,
		Message: "Something wrong with your input data",
	}
	return c.Status(500).JSON(response)
}

func ResponseCityCountryNotFound(c *fiber.Ctx, country *entities.Country, err error) error {
	if err != nil || country.ID == "" {
		response := ResponseCityStructure{
			Error:   err,
			Message: "Country not found",
		}
		return c.Status(404).JSON(response)
	}
	return nil
}

// Ответы

func ResponseCities(c *fiber.Ctx, cities []*entities.City, err error) error {

	data := ResponseDataCity{Cities: cities}
	var statusCode int
	var message string
	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else if len(cities) == 0 {
		statusCode = 404
		message = "Cities not found"
	} else {
		statusCode = 201
		message = "Cities found"
	}

	response := ResponseCityStructure{
		Data:    data,
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseCity(c *fiber.Ctx, city *entities.City, err error) error {

	cities := make([]*entities.City, 0)
	if city.ID != "" {
		cities = append(cities, city)
	}
	return ResponseCities(c, cities, err)

}

func ResponseCityCreate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "City has created"
	}
	response := ResponseCityStructure{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseCityUpdate(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "City has updated"
	}
	response := ResponseCityStructure{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}

func ResponseCityDelete(c *fiber.Ctx, err error) error {

	var statusCode int
	var message string

	if err != nil {
		statusCode = 500
		message = "Unexpected error"
	} else {
		statusCode = 201
		message = "City has deleted"
	}
	response := ResponseCityStructure{
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)

}
