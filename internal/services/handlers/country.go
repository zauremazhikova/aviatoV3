package handlers

import (
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transport"
	"github.com/gofiber/fiber/v2"
)

func GetAllCountries(c *fiber.Ctx) error {

	responseCountry, err := repositories.GetCountries()
	return transport.ResponseCountries(c, responseCountry, err)

}

func GetSingleCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	country, err := repositories.GetCountry(id)
	return transport.ResponseCountry(c, country, err)

}

func CreateCountry(c *fiber.Ctx) error {

	country, err := transport.ValidateCountryInsertData(c)
	if err != nil {
		return err
	}
	err = repositories.CreateCountry(country)
	return transport.ResponseCountryCreate(c, err)

}

func UpdateCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	country, err := repositories.GetCountry(id)

	err = transport.ResponseCountryNotFound(c, country, err)
	if err != nil {
		return err
	}

	updateCountryData, err := transport.ValidateCountryUpdateData(c)
	if err != nil {
		return err
	}

	country.Name = updateCountryData.Name
	err = repositories.UpdateCountry(country)

	return transport.ResponseCountryUpdate(c, err)
}

func DeleteCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	country, err := repositories.GetCountry(id)

	err = transport.ResponseCountryNotFound(c, country, err)
	if err != nil {
		return err
	}

	err = repositories.DeleteCountry(id)
	return transport.ResponseCountryDelete(c, err)

}
