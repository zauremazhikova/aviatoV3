package entityHandlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
)

// Validating structure

type UpdateCountryStructure struct {
	Name string `json:"name"`
}

type InsertCountryStructure struct {
	Name string `json:"name"`
}

// Methods

func GetAllCountries() (a []*entities.Country, err error) {

	responseCountry, err := repositories.GetCountries()
	return responseCountry, err

}

func GetSingleCountry(id string) (a *entities.Country, err error) {

	country, err := repositories.GetCountry(id)
	return country, err

}

func CreateCountry(insertStruct *InsertCountryStructure) error {

	country := new(entities.Country)
	country.Name = insertStruct.Name
	err := repositories.CreateCountry(country)
	return err

}

func UpdateCountry(country *entities.Country, updateCountryData *UpdateCountryStructure) error {

	country.Name = updateCountryData.Name
	err := repositories.UpdateCountry(country)
	return err

}

func DeleteCountry(id string) error {

	err := repositories.DeleteCountry(id)
	return err

}
