package transport

import (
	"aviatoV3/internal/entities"
	"github.com/gofiber/fiber/v2"
)

type ResponseStructureCity struct {
	StatusCode int              `json:"statusCode"`
	Data       ResponseDataCity `json:"data"`
	Error      error            `json:"error"`
	Message    string           `json:"message"`
}

type ResponseDataCity struct {
	Cities []*entities.City `json:"cities"`
}

type UpdateCity struct {
	Name      string `json:"name"`
	CountryID string `json:"countryID"`
}

// Валидация входящих данных

func ValidateCityInsertData(c *fiber.Ctx) (*entities.City, error) {
	city := new(entities.City)
	err := c.BodyParser(city)

	if err != nil {
		return city, c.JSON(ResponseCityInputError(err))
	}
	return city, nil
}

func ValidateCityUpdateData(c *fiber.Ctx) (*UpdateCity, error) {

	var updateCityData UpdateCity
	err := c.BodyParser(&updateCityData)

	if err != nil {
		return &updateCityData, c.JSON(ResponseCityInputError(err))
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

func ResponseCityInputError(err error) ResponseStructureCity {

	response := ResponseStructureCity{
		StatusCode: 500,
		Data:       ResponseDataCity{},
		Error:      err,
		Message:    "Something wrong with your input data",
	}
	return response
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

	response := ResponseStructureCity{
		StatusCode: statusCode,
		Data:       data,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

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
	response := ResponseStructureCity{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

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
	response := ResponseStructureCity{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

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
	response := ResponseStructureCity{
		StatusCode: statusCode,
		Error:      err,
		Message:    message,
	}
	return c.JSON(response)

}
