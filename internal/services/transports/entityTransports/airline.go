package entityTransports

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/services/handlers/entityHandlers"
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

// Response making

func ResponseAirlines(c *fiber.Ctx, airlines []*entities.Airline, err error, statusCode int, message string) error {

	response := ResponseStructureAirline{
		Data:    ResponseDataAirline{Airlines: airlines},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Methods

func GetAllAirlines(c *fiber.Ctx) error {

	airlines, err := entityHandlers.GetAllAirlines()
	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Unexpected error")
	} else if len(airlines) == 0 {
		return ResponseAirlines(c, airlines, err, 404, "Airlines not found")
	}
	return ResponseAirlines(c, airlines, err, 201, "Airlines found")
}

func GetSingleAirline(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := entityHandlers.GetSingleAirline(id)

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

	var insertStructure entityHandlers.InsertAirlineStructure
	err := c.BodyParser(&insertStructure)
	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Something wrong with your input data")
	}

	err = entityHandlers.CreateAirline(&insertStructure)
	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Unexpected error")
	}

	return ResponseAirlines(c, airlines, err, 201, "Airline has created")

}

func UpdateAirline(c *fiber.Ctx) error {

	airlines := make([]*entities.Airline, 0)

	var updateAirlineData entityHandlers.UpdateAirlineStructure
	err := c.BodyParser(&updateAirlineData)

	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Something wrong with your input data")
	}

	id := c.Params("id")
	airline, err := entityHandlers.GetSingleAirline(id)

	if airline.ID == "" {
		return ResponseAirlines(c, airlines, err, 404, "Airline not found")
	}
	err = entityHandlers.UpdateAirline(airline, &updateAirlineData)

	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Unexpected error")
	}
	return ResponseAirlines(c, airlines, err, 201, "Airline has updated")

}

func DeleteAirline(c *fiber.Ctx) error {

	airlines := make([]*entities.Airline, 0)

	id := c.Params("id")
	airline, err := entityHandlers.GetSingleAirline(id)

	if airline.ID == "" {
		return ResponseAirlines(c, airlines, err, 404, "Airline not found")
	}

	err = entityHandlers.DeleteAirline(id)
	if err != nil {
		return ResponseAirlines(c, airlines, err, 500, "Unexpected error")
	}

	return ResponseAirlines(c, airlines, err, 201, "Airline has deleted")

}
