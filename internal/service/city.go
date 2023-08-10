package service

import (
	"aviatoV3/internal/entity"
	"aviatoV3/internal/repository"
	"github.com/gofiber/fiber/v2"
)

// Response structure

type CityResponseStructure struct {
	Data    CityResponseData `json:"data"`
	Error   error            `json:"error"`
	Message string           `json:"message"`
}

type CityResponseData struct {
	Cities []*entity.City `json:"cities"`
}

// Validation structures

type CityUpdateStructure struct {
	Name string `json:"name"`
}

type CityInsertStructure struct {
	Name string `json:"name"`
}

// Response making

func CityResponse(c *fiber.Ctx, cities []*entity.City, err error, statusCode int, message string) error {

	response := CityResponseStructure{
		Data:    CityResponseData{Cities: cities},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Validation

func CityCreateValidation(c *fiber.Ctx, cities []*entity.City) (*CityInsertStructure, error) {

	var insertStructure CityInsertStructure
	err := c.BodyParser(&insertStructure)

	if err != nil {
		err = CityResponse(c, cities, err, 500, "Something wrong with your input data")
	}
	return &insertStructure, err

}

func CityUpdateValidation(c *fiber.Ctx, cities []*entity.City) (*CityUpdateStructure, error) {

	var updateStructure CityUpdateStructure
	err := c.BodyParser(&updateStructure)

	if err != nil {
		err = CityResponse(c, cities, err, 500, "Something wrong with your input data")
	}
	return &updateStructure, err

}

// Methods

func GetAllCities(c *fiber.Ctx) error {

	cities, err := repository.GetCities()
	if err != nil {
		return CityResponse(c, cities, err, 500, "Unexpected error")
	} else if len(cities) == 0 {
		return CityResponse(c, cities, err, 404, "Cities not found")
	}
	return CityResponse(c, cities, err, 201, "Cities found")
}

func GetSingleCity(c *fiber.Ctx) error {

	id := c.Params("id")
	city, err := repository.GetCity(id)

	cities := make([]*entity.City, 0)
	if city.ID != "" {
		cities = append(cities, city)
	}

	if err != nil {
		return CityResponse(c, cities, err, 500, "Unexpected error")
	} else if len(cities) == 0 {
		return CityResponse(c, cities, err, 404, "City not found")
	}

	return CityResponse(c, cities, err, 201, "City found")

}

func CreateCity(c *fiber.Ctx) error {

	cities := make([]*entity.City, 0)
	insertStructure, err := CityCreateValidation(c, cities)
	if err != nil {
		return err
	}

	city := new(entity.City)
	city.Name = insertStructure.Name
	err = repository.CreateCity(city)
	if err != nil {
		return CityResponse(c, cities, err, 500, "Unexpected error")
	}
	cities = append(cities, city)
	return CityResponse(c, cities, err, 201, "City has created")

}

func UpdateCity(c *fiber.Ctx) error {

	cities := make([]*entity.City, 0)
	updateStructure, err := CityUpdateValidation(c, cities)
	if err != nil {
		return err
	}

	id := c.Params("id")
	city, err := repository.GetCity(id)
	if city.ID == "" {
		return CityResponse(c, cities, err, 404, "City not found")
	}

	city.Name = updateStructure.Name
	err = repository.UpdateCity(city)
	if err != nil {
		return CityResponse(c, cities, err, 500, "Unexpected error")
	}
	cities = append(cities, city)
	return CityResponse(c, cities, err, 201, "City has updated")

}

func DeleteCity(c *fiber.Ctx) error {

	cities := make([]*entity.City, 0)

	id := c.Params("id")
	city, err := repository.GetCity(id)

	if city.ID == "" {
		return CityResponse(c, cities, err, 404, "City not found")
	}

	err = repository.DeleteCity(id)
	if err != nil {
		return CityResponse(c, cities, err, 500, "Unexpected error")
	}

	return CityResponse(c, cities, err, 201, "City has deleted")

}
