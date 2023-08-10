package entityTransports

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/services/handlers/entityHandlers"
	"github.com/gofiber/fiber/v2"
)

// Response structure

type ResponseStructureCountry struct {
	Data    ResponseDataCountry `json:"data"`
	Error   error               `json:"error"`
	Message string              `json:"message"`
}

type ResponseDataCountry struct {
	Countries []*entities.Country `json:"countries"`
}

// Response making

func ResponseCountries(c *fiber.Ctx, countries []*entities.Country, err error, statusCode int, message string) error {

	response := ResponseStructureCountry{
		Data:    ResponseDataCountry{Countries: countries},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Methods

func GetAllCountries(c *fiber.Ctx) error {

	countries, err := entityHandlers.GetAllCountries()
	if err != nil {
		return ResponseCountries(c, countries, err, 500, "Unexpected error")
	} else if len(countries) == 0 {
		return ResponseCountries(c, countries, err, 404, "Countries not found")
	}
	return ResponseCountries(c, countries, err, 201, "Countries found")
}

func GetSingleCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	country, err := entityHandlers.GetSingleCountry(id)

	countries := make([]*entities.Country, 0)
	if country.ID != "" {
		countries = append(countries, country)
	}

	if err != nil {
		return ResponseCountries(c, countries, err, 500, "Unexpected error")
	} else if len(countries) == 0 {
		return ResponseCountries(c, countries, err, 404, "Country not found")
	}

	return ResponseCountries(c, countries, err, 201, "Country found")

}

func CreateCountry(c *fiber.Ctx) error {

	countries := make([]*entities.Country, 0)

	var insertStructure entityHandlers.InsertCountryStructure
	err := c.BodyParser(&insertStructure)
	if err != nil {
		return ResponseCountries(c, countries, err, 500, "Something wrong with your input data")
	}

	err = entityHandlers.CreateCountry(&insertStructure)
	if err != nil {
		return ResponseCountries(c, countries, err, 500, "Unexpected error")
	}

	return ResponseCountries(c, countries, err, 201, "Country has created")

}

func UpdateCountry(c *fiber.Ctx) error {

	var updateCountryData entityHandlers.UpdateCountryStructure
	countries := make([]*entities.Country, 0)

	err := c.BodyParser(&updateCountryData)
	if err != nil {
		return ResponseCountries(c, countries, err, 500, "Something wrong with your input data")
	}

	id := c.Params("id")
	country, err := entityHandlers.GetSingleCountry(id)
	if country.ID == "" {
		return ResponseCountries(c, countries, err, 404, "Country not found")
	}

	err = entityHandlers.UpdateCountry(country, &updateCountryData)
	if err != nil {
		return ResponseCountries(c, countries, err, 500, "Unexpected error")
	}
	return ResponseCountries(c, countries, err, 201, "Country has updated")

}

func DeleteCountry(c *fiber.Ctx) error {

	countries := make([]*entities.Country, 0)

	id := c.Params("id")
	country, err := entityHandlers.GetSingleCountry(id)
	if country.ID == "" {
		return ResponseCountries(c, countries, err, 404, "Country not found")
	}

	err = entityHandlers.DeleteCountry(id)
	if err != nil {
		return ResponseCountries(c, countries, err, 500, "Unexpected error")
	}

	return ResponseCountries(c, countries, err, 201, "Country has deleted")

}
