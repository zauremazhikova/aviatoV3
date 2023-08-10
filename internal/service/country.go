package service

import (
	"aviatoV3/internal/entity"
	"aviatoV3/internal/repository"
	"github.com/gofiber/fiber/v2"
)

// Response structure

type CountryResponseStructure struct {
	Data    CountryResponseData `json:"data"`
	Error   error               `json:"error"`
	Message string              `json:"message"`
}

type CountryResponseData struct {
	Countries []*entity.Country `json:"countries"`
}

// Validation structures

type CountryUpdateStructure struct {
	Name string `json:"name"`
}

type CountryInsertStructure struct {
	Name string `json:"name"`
}

// Response making

func CountryResponse(c *fiber.Ctx, countries []*entity.Country, err error, statusCode int, message string) error {

	response := CountryResponseStructure{
		Data:    CountryResponseData{Countries: countries},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Validation

func CountryCreateValidation(c *fiber.Ctx, countries []*entity.Country) (*CountryInsertStructure, error) {

	var insertStructure CountryInsertStructure
	err := c.BodyParser(&insertStructure)

	if err != nil {
		err = CountryResponse(c, countries, err, 500, "Something wrong with your input data")
	}
	return &insertStructure, err

}

func CountryUpdateValidation(c *fiber.Ctx, countries []*entity.Country) (*CountryUpdateStructure, error) {

	var updateStructure CountryUpdateStructure
	err := c.BodyParser(&updateStructure)

	if err != nil {
		err = CountryResponse(c, countries, err, 500, "Something wrong with your input data")
	}
	return &updateStructure, err

}

// Methods

func GetAllCountries(c *fiber.Ctx) error {

	countries, err := repository.GetCountries()
	if err != nil {
		return CountryResponse(c, countries, err, 500, "Unexpected error")
	} else if len(countries) == 0 {
		return CountryResponse(c, countries, err, 404, "Countries not found")
	}
	return CountryResponse(c, countries, err, 201, "Countries found")
}

func GetSingleCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	country, err := repository.GetCountry(id)

	countries := make([]*entity.Country, 0)
	if country.ID != "" {
		countries = append(countries, country)
	}

	if err != nil {
		return CountryResponse(c, countries, err, 500, "Unexpected error")
	} else if len(countries) == 0 {
		return CountryResponse(c, countries, err, 404, "Country not found")
	}

	return CountryResponse(c, countries, err, 201, "Country found")

}

func CreateCountry(c *fiber.Ctx) error {

	countries := make([]*entity.Country, 0)
	insertStructure, err := CountryCreateValidation(c, countries)
	if err != nil {
		return err
	}

	country := new(entity.Country)
	country.Name = insertStructure.Name
	err = repository.CreateCountry(country)
	if err != nil {
		return CountryResponse(c, countries, err, 500, "Unexpected error")
	}
	countries = append(countries, country)
	return CountryResponse(c, countries, err, 201, "Country has created")

}

func UpdateCountry(c *fiber.Ctx) error {

	countries := make([]*entity.Country, 0)
	updateStructure, err := CountryUpdateValidation(c, countries)
	if err != nil {
		return err
	}

	id := c.Params("id")
	country, err := repository.GetCountry(id)
	if country.ID == "" {
		return CountryResponse(c, countries, err, 404, "Country not found")
	}

	country.Name = updateStructure.Name
	err = repository.UpdateCountry(country)
	if err != nil {
		return CountryResponse(c, countries, err, 500, "Unexpected error")
	}
	countries = append(countries, country)
	return CountryResponse(c, countries, err, 201, "Country has updated")

}

func DeleteCountry(c *fiber.Ctx) error {

	countries := make([]*entity.Country, 0)

	id := c.Params("id")
	country, err := repository.GetCountry(id)

	if country.ID == "" {
		return CountryResponse(c, countries, err, 404, "Country not found")
	}

	err = repository.DeleteCountry(id)
	if err != nil {
		return CountryResponse(c, countries, err, 500, "Unexpected error")
	}

	return CountryResponse(c, countries, err, 201, "Country has deleted")

}
