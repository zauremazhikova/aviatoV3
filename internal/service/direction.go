package service

import (
	"aviatoV3/internal/entity"
	"aviatoV3/internal/repository"
	"github.com/gofiber/fiber/v2"
)

// Response structure

type DirectionResponseStructure struct {
	Data    DirectionResponseData `json:"data"`
	Error   error                 `json:"error"`
	Message string                `json:"message"`
}

type DirectionResponseData struct {
	Directions []*entity.Direction `json:"directions"`
}

// Validation structures

type DirectionUpdateStructure struct {
	OriginCityID      string `json:"originCityID"`
	DestinationCityID string `json:"destinationCityID"`
	AirlineID         string `json:"airlineID"`
}

type DirectionInsertStructure struct {
	OriginCityID      string `json:"originCityID"`
	DestinationCityID string `json:"destinationCityID"`
	AirlineID         string `json:"airlineID"`
}

// Response making

func DirectionResponse(c *fiber.Ctx, directions []*entity.Direction, err error, statusCode int, message string) error {

	response := DirectionResponseStructure{
		Data:    DirectionResponseData{Directions: directions},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Validation

func DirectionCreateValidation(c *fiber.Ctx, directions []*entity.Direction) (*DirectionInsertStructure, error) {

	var insertStructure DirectionInsertStructure
	err := c.BodyParser(&insertStructure)

	if err != nil {
		err = DirectionResponse(c, directions, err, 500, "Something wrong with your input data")
	}
	return &insertStructure, err

}

func DirectionUpdateValidation(c *fiber.Ctx, cities []*entity.Direction) (*DirectionUpdateStructure, error) {

	var updateStructure DirectionUpdateStructure
	err := c.BodyParser(&updateStructure)

	if err != nil {
		err = DirectionResponse(c, cities, err, 500, "Something wrong with your input data")
	}
	return &updateStructure, err

}

// Methods

func GetAllDirections(c *fiber.Ctx) error {

	directions, err := repository.GetDirections()
	if err != nil {
		return DirectionResponse(c, directions, err, 500, "Unexpected error")
	} else if len(directions) == 0 {
		return DirectionResponse(c, directions, err, 404, "Directions not found")
	}
	return DirectionResponse(c, directions, err, 201, "Directions found")
}

func GetSingleDirection(c *fiber.Ctx) error {

	id := c.Params("id")
	direction, err := repository.GetDirection(id)

	directions := make([]*entity.Direction, 0)
	if direction.ID != "" {
		directions = append(directions, direction)
	}

	if err != nil {
		return DirectionResponse(c, directions, err, 500, "Unexpected error")
	} else if len(directions) == 0 {
		return DirectionResponse(c, directions, err, 404, "Direction not found")
	}

	return DirectionResponse(c, directions, err, 201, "Direction found")

}

func CreateDirection(c *fiber.Ctx) error {

	directions := make([]*entity.Direction, 0)
	insertStructure, err := DirectionCreateValidation(c, directions)
	if err != nil {
		return err
	}

	originCity, err := repository.GetCity(insertStructure.OriginCityID)
	if err != nil || originCity.ID == "" {
		return DirectionResponse(c, directions, err, 404, "Origin city not found")
	}

	destinationCity, err := repository.GetCity(insertStructure.DestinationCityID)
	if err != nil || destinationCity.ID == "" {
		return DirectionResponse(c, directions, err, 404, "Destination city not found")
	}
	airline, err := repository.GetAirline(insertStructure.AirlineID)
	if err != nil || airline.ID == "" {
		return DirectionResponse(c, directions, err, 404, "Airlines not found")
	}

	direction := new(entity.Direction)
	direction.OriginCity = *originCity
	direction.DestinationCity = *destinationCity
	direction.Airline = *airline

	err = repository.CreateDirection(direction)
	if err != nil {
		return DirectionResponse(c, directions, err, 500, "Unexpected error")
	}
	directions = append(directions, direction)
	return DirectionResponse(c, directions, err, 201, "Direction has created")

}

func UpdateDirection(c *fiber.Ctx) error {

	directions := make([]*entity.Direction, 0)
	updateStructure, err := DirectionUpdateValidation(c, directions)
	if err != nil {
		return err
	}

	id := c.Params("id")
	direction, err := repository.GetDirection(id)
	if direction.ID == "" {
		return DirectionResponse(c, directions, err, 404, "Direction not found")
	}

	originCity, err := repository.GetCity(updateStructure.OriginCityID)
	if err != nil || originCity.ID == "" {
		return DirectionResponse(c, directions, err, 404, "Origin city not found")
	}

	destinationCity, err := repository.GetCity(updateStructure.DestinationCityID)
	if err != nil || destinationCity.ID == "" {
		return DirectionResponse(c, directions, err, 404, "Destination city not found")
	}
	airline, err := repository.GetAirline(updateStructure.AirlineID)
	if err != nil || airline.ID == "" {
		return DirectionResponse(c, directions, err, 404, "Airlines not found")
	}

	direction.OriginCity = *originCity
	direction.DestinationCity = *destinationCity
	direction.Airline = *airline

	err = repository.UpdateDirection(direction)
	if err != nil {
		return DirectionResponse(c, directions, err, 500, "Unexpected error")
	}
	directions = append(directions, direction)
	return DirectionResponse(c, directions, err, 201, "Direction has updated")

}

func DeleteDirection(c *fiber.Ctx) error {

	directions := make([]*entity.Direction, 0)

	id := c.Params("id")
	direction, err := repository.GetDirection(id)

	if direction.ID == "" {
		return DirectionResponse(c, directions, err, 404, "Direction not found")
	}

	err = repository.DeleteDirection(id)
	if err != nil {
		return DirectionResponse(c, directions, err, 500, "Unexpected error")
	}

	return DirectionResponse(c, directions, err, 201, "Direction has deleted")

}
