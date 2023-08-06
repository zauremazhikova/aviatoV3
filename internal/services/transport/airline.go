package transport

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int   `json:"statusCode"`
	Data       Data  `json:"data"`
	Error      error `json:"error"`
}

type Data struct {
	Airlines []*entities.Airline `json:"airlines"`
}

type UpdateAirline struct {
	Name string `json:"name"`
}

func ResponseAirlines(c *fiber.Ctx, airlines []*entities.Airline, err error) error {

	data := Data{Airlines: airlines}
	var statusCode int
	if err != nil {
		statusCode = 500
	} else {
		statusCode = 201
	}

	response := Response{
		StatusCode: statusCode,
		Data:       data,
		Error:      err,
	}
	return c.JSON(response)

}

func ResponseAirline(c *fiber.Ctx, airline *entities.Airline, err error) error {

	airlines := make([]*entities.Airline, 0)
	if airline.ID != "" {
		airlines = append(airlines, airline)
	}
	data := Data{Airlines: airlines}

	var statusCode int
	if err != nil {
		statusCode = 500
	} else {
		statusCode = 201
	}

	response := Response{
		StatusCode: statusCode,
		Data:       data,
		Error:      err,
	}
	return c.JSON(response)

}

func ValidateInsertData(c *fiber.Ctx) (*entities.Airline, error) {

	airline := new(entities.Airline)
	err := c.BodyParser(airline)

	if err != nil {
		response := Response{
			StatusCode: 500,
			Data:       Data{},
			Error:      err,
		}
		return airline, c.JSON(response)
	}
	return airline, nil

}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": err})
	}

	var updateAirlineData UpdateAirline
	err = c.BodyParser(&updateAirlineData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	airline.Name = updateAirlineData.Name
	err = repositories.UpdateAirline(airline)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Airline has not updated", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Airline has Updated", "data": airline})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := repositories.GetAirline(id)

	if airline.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": nil})
	}

	err = repositories.DeleteAirline(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete Airline", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Airline deleted"})
}
