package entityHandlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
)

// Validating structure

type UpdateAirlineStructure struct {
	Name string `json:"name"`
}

type InsertAirlineStructure struct {
	Name string `json:"name"`
}

// Methods

func GetAllAirlines() (a []*entities.Airline, err error) {

	responseAirline, err := repositories.GetAirlines()
	return responseAirline, err

}

func GetSingleAirline(id string) (a *entities.Airline, err error) {

	airline, err := repositories.GetAirline(id)
	return airline, err

}

func CreateAirline(insertStruct *InsertAirlineStructure) error {

	airline := new(entities.Airline)
	airline.Name = insertStruct.Name
	err := repositories.CreateAirline(airline)
	return err

}

func UpdateAirline(airline *entities.Airline, updateAirlineData *UpdateAirlineStructure) error {

	airline.Name = updateAirlineData.Name
	err := repositories.UpdateAirline(airline)
	return err

}

func DeleteAirline(id string) error {

	err := repositories.DeleteAirline(id)
	return err

}
