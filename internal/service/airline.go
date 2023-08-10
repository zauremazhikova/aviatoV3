package service

import (
	"aviatoV3/internal/entity"
	"aviatoV3/internal/repository"
	"github.com/gofiber/fiber/v2"
)

// Response structure

type AirlineResponseStructure struct {
	Data    AirlineResponseData `json:"data"`
	Error   error               `json:"error"`
	Message string              `json:"message"`
}

type AirlineResponseData struct {
	Airlines []*entity.Airline `json:"airlines"`
}

// Validation structures

type AirlineUpdateStructure struct {
	Name string `json:"name"`
}

type AirlineInsertStructure struct {
	Name string `json:"name"`
}

// Response making

func AirlineResponse(c *fiber.Ctx, airlines []*entity.Airline, err error, statusCode int, message string) error {

	response := AirlineResponseStructure{
		Data:    AirlineResponseData{Airlines: airlines},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Validation

func AirlineCreateValidation(c *fiber.Ctx, airlines []*entity.Airline) (*AirlineInsertStructure, error) {

	var insertStructure AirlineInsertStructure
	err := c.BodyParser(&insertStructure)

	if err != nil {
		err = AirlineResponse(c, airlines, err, 500, "Something wrong with your input data")
	}
	return &insertStructure, err

}

func AirlineUpdateValidation(c *fiber.Ctx, airlines []*entity.Airline) (*AirlineUpdateStructure, error) {

	var updateStructure AirlineUpdateStructure
	err := c.BodyParser(&updateStructure)

	if err != nil {
		err = AirlineResponse(c, airlines, err, 500, "Something wrong with your input data")
	}
	return &updateStructure, err

}

// Methods

func GetAllAirlines(c *fiber.Ctx) error {

	airlines, err := repository.GetAirlines()
	if err != nil {
		return AirlineResponse(c, airlines, err, 500, "Unexpected error")
	} else if len(airlines) == 0 {
		return AirlineResponse(c, airlines, err, 404, "Airlines not found")
	}
	return AirlineResponse(c, airlines, err, 201, "Airlines found")
}

func GetSingleAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repository.GetAirline(id)

	airlines := make([]*entity.Airline, 0)
	if airline.ID != "" {
		airlines = append(airlines, airline)
	}

	if err != nil {
		return AirlineResponse(c, airlines, err, 500, "Unexpected error")
	} else if len(airlines) == 0 {
		return AirlineResponse(c, airlines, err, 404, "Airline not found")
	}

	return AirlineResponse(c, airlines, err, 201, "Airline found")

}

func CreateAirline(c *fiber.Ctx) error {

	airlines := make([]*entity.Airline, 0)
	insertStructure, err := AirlineCreateValidation(c, airlines)
	if err != nil {
		return err
	}

	airline := new(entity.Airline)
	airline.Name = insertStructure.Name
	err = repository.CreateAirline(airline)
	if err != nil {
		return AirlineResponse(c, airlines, err, 500, "Unexpected error")
	}
	airlines = append(airlines, airline)
	return AirlineResponse(c, airlines, err, 201, "Airline has created")

}

func UpdateAirline(c *fiber.Ctx) error {

	airlines := make([]*entity.Airline, 0)
	updateStructure, err := AirlineUpdateValidation(c, airlines)
	if err != nil {
		return err
	}

	id := c.Params("id")
	airline, err := repository.GetAirline(id)
	if airline.ID == "" {
		return AirlineResponse(c, airlines, err, 404, "Airline not found")
	}

	airline.Name = updateStructure.Name
	err = repository.UpdateAirline(airline)
	if err != nil {
		return AirlineResponse(c, airlines, err, 500, "Unexpected error")
	}
	airlines = append(airlines, airline)
	return AirlineResponse(c, airlines, err, 201, "Airline has updated")

}

func DeleteAirline(c *fiber.Ctx) error {

	airlines := make([]*entity.Airline, 0)

	id := c.Params("id")
	airline, err := repository.GetAirline(id)

	if airline.ID == "" {
		return AirlineResponse(c, airlines, err, 404, "Airline not found")
	}

	err = repository.DeleteAirline(id)
	if err != nil {
		return AirlineResponse(c, airlines, err, 500, "Unexpected error")
	}

	return AirlineResponse(c, airlines, err, 201, "Airline has deleted")

}
