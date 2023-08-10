package entityHandlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

// Response structure

type ResponseStructureAirline struct {
	Data    ResponseDataAirline `json:"data"`
	Error   error               `json:"error"`
	Message string              `json:"message"`
}

type ResponseDataAirline struct {
	Airlines []*entities.Airline `json:"airlines"`
}

// Validation structures

type AirlineUpdateStructure struct {
	Name string `json:"name"`
}

type AirlineInsertStructure struct {
	Name string `json:"name"`
}

// Response making

func ResponseAirlines(c *fiber.Ctx, airlines []*entities.Airline, err error, statusCode int, message string) error {

	response := ResponseStructureAirline{
		Data:    ResponseDataAirline{Airlines: airlines},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Validation

func AirlineCreateValidation(c *fiber.Ctx, airlines []*entities.Airline) (*AirlineInsertStructure, error) {

	var insertStructure AirlineInsertStructure
	err := c.BodyParser(&insertStructure)

	if err != nil {
		err = ResponseAirlines(c, airlines, err, 500, "Something wrong with your input data")
	}
	return &insertStructure, err

}

func AirlineUpdateValidation(c *fiber.Ctx, airlines []*entities.Airline) (*AirlineUpdateStructure, error) {

	var updateAirlineData AirlineUpdateStructure
	err := c.BodyParser(&updateAirlineData)

	if err != nil {
		err = ResponseAirlines(c, airlines, err, 500, "Something wrong with your input data")
	}
	return &updateAirlineData, err

}

// Methods

func GetAllAirlines(c *fiber.Ctx) error {

	airlines, err := repositories.GetAirlines()
	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Unexpected error")
	} else if len(airlines) == 0 {
		return ResponseAirlines(c, airlines, err, 404, "Airlines not found")
	}
	return ResponseAirlines(c, airlines, err, 201, "Airlines found")
}

func GetSingleAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	airlines := make([]*entities.Airline, 0)
	if airline.ID != "" {
		airlines = append(airlines, airline)
	}

	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Unexpected error")
	} else if len(airlines) == 0 {
		return ResponseAirlines(c, airlines, err, 404, "Airline not found")
	}

	return ResponseAirlines(c, airlines, err, 201, "Airline found")

}

func CreateAirline(c *fiber.Ctx) error {

	airlines := make([]*entities.Airline, 0)
	insertStructure, err := AirlineCreateValidation(c, airlines)
	if err != nil {
		return err
	}

	airline := new(entities.Airline)
	airline.Name = insertStructure.Name
	err = repositories.CreateAirline(airline)
	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Unexpected error")
	}
	airlines = append(airlines, airline)
	return ResponseAirlines(c, airlines, err, 201, "Airline has created")

}

func UpdateAirline(c *fiber.Ctx) error {

	airlines := make([]*entities.Airline, 0)
	updateStructure, err := AirlineUpdateValidation(c, airlines)
	if err != nil {
		return err
	}

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)
	if airline.ID == "" {
		return ResponseAirlines(c, airlines, err, 404, "Airline not found")
	}

	airline.Name = updateStructure.Name
	err = repositories.UpdateAirline(airline)
	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Unexpected error")
	}
	airlines = append(airlines, airline)
	return ResponseAirlines(c, airlines, err, 201, "Airline has updated")

}

func DeleteAirline(c *fiber.Ctx) error {

	airlines := make([]*entities.Airline, 0)

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	if airline.ID == "" {
		return ResponseAirlines(c, airlines, err, 404, "Airline not found")
	}

	err = repositories.DeleteAirline(id)
	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Unexpected error")
	}

	return ResponseAirlines(c, airlines, err, 201, "Airline has deleted")

}
