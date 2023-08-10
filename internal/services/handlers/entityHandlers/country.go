package entityHandlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transports/entityTransports"
	"github.com/gofiber/fiber/v2"
)

func GetAllCountries(c *fiber.Ctx) error {

	responseCountry, err := repositories.GetCountries()
	return entityTransports.ResponseCountries(c, responseCountry, err)

}

func GetSingleCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	country, err := repositories.GetCountry(id)
	return entityTransports.ResponseCountry(c, country, err)

}

func CreateCountry(c *fiber.Ctx) error {

	insertStruct, err := entityTransports.ValidateCountryInsertData(c)
	if err != nil {
		return err
	}
	country := new(entities.Country)
	country.Name = insertStruct.Name

	err = repositories.CreateCountry(country)
	return entityTransports.ResponseCountryCreate(c, err)

}

func UpdateCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	country, err := repositories.GetCountry(id)

	err = entityTransports.ResponseCountryNotFound(c, country, err)
	if err != nil {
		return err
	}

	updateCountryData, err := entityTransports.ValidateCountryUpdateData(c)
	if err != nil {
		return err
	}

	country.Name = updateCountryData.Name
	err = repositories.UpdateCountry(country)

	return entityTransports.ResponseCountryUpdate(c, err)
}

func DeleteCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	country, err := repositories.GetCountry(id)

	err = entityTransports.ResponseCountryNotFound(c, country, err)
	if err != nil {
		return err
	}

	err = repositories.DeleteCountry(id)
	return entityTransports.ResponseCountryDelete(c, err)

}
