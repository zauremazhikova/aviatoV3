package entityHandlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transports/entityTransports"
	"github.com/gofiber/fiber/v2"
)

func GetAllCities(c *fiber.Ctx) error {

	responseCities, err := repositories.GetCities()
	return entityTransports.ResponseCities(c, responseCities, err)

}

func GetSingleCity(c *fiber.Ctx) error {

	id := c.Params("id")
	city, err := repositories.GetCity(id)
	return entityTransports.ResponseCity(c, city, err)

}

func CreateCity(c *fiber.Ctx) error {

	insertStruct, err := entityTransports.ValidateCityInsertData(c)
	if err != nil {
		return err
	}

	currentCountry, err := repositories.GetCountry(insertStruct.CountryID)
	err = entityTransports.ResponseCityCountryNotFound(c, currentCountry, err)
	if err != nil {
		return err
	}

	city := new(entities.City)
	city.Name = insertStruct.Name
	city.Country = *currentCountry

	err = repositories.CreateCity(city)
	return entityTransports.ResponseCityCreate(c, err)

}

func UpdateCity(c *fiber.Ctx) error {

	id := c.Params("id")
	city, err := repositories.GetCity(id)

	err = entityTransports.ResponseCityNotFound(c, city, err)
	if err != nil {
		return err
	}

	updateCityData, err := entityTransports.ValidateCityUpdateData(c)
	if err != nil {
		return err
	}

	currentCountry, err := repositories.GetCountry(updateCityData.CountryID)
	err = entityTransports.ResponseCityCountryNotFound(c, currentCountry, err)
	if err != nil {
		return err
	}

	city.Name = updateCityData.Name
	city.Country = *currentCountry
	err = repositories.UpdateCity(city)

	return entityTransports.ResponseCityUpdate(c, err)
}

func DeleteCity(c *fiber.Ctx) error {

	id := c.Params("id")
	city, err := repositories.GetCity(id)

	err = entityTransports.ResponseCityNotFound(c, city, err)
	if err != nil {
		return err
	}

	err = repositories.DeleteCity(id)
	return entityTransports.ResponseCityDelete(c, err)

}
