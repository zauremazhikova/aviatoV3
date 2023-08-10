package service

import (
	"aviatoV3/config"
	"aviatoV3/internal/entity"
	"aviatoV3/internal/repository"
	"github.com/gofiber/fiber/v2"
	"slices"
	"time"
)

var flightsMap [][]*entity.Flight

type FlightVariantsSearchStruct struct {
	OriginCityID      string    `json:"OriginCityID"`
	DestinationCityID string    `json:"DestinationCityID"`
	DepartureTime     time.Time `json:"departureTime"`
}

type FlightVariantsResponseStructure struct {
	Data    FlightVariantsResponseData `json:"data"`
	Error   error                      `json:"error"`
	Message string                     `json:"message"`
}

type FlightVariantsResponseData struct {
	Flights [][]*entity.Flight `json:"Flights"`
}

// Response making

func FlightVariantsResponse(c *fiber.Ctx, flights [][]*entity.Flight, err error, statusCode int, message string) error {

	response := FlightVariantsResponseStructure{
		Data:    FlightVariantsResponseData{Flights: flights},
		Error:   err,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

// Validation

func FlightVariantsValidation(c *fiber.Ctx, flights [][]*entity.Flight) (*FlightVariantsSearchStruct, error) {

	var searchStructure FlightVariantsSearchStruct
	err := c.BodyParser(&searchStructure)

	if err != nil {
		err = FlightVariantsResponse(c, flights, err, 500, "Something wrong with your input data")
	}
	return &searchStructure, err

}

// Flight searching

func GetFlightsByOriginAndDestination(c *fiber.Ctx) error {

	flightsMap = make([][]*entity.Flight, 0)

	searchData, err := FlightVariantsValidation(c, flightsMap)
	if err != nil {
		return err
	}
	// maxStop - это максимальное количество пересадок. Настраивается в config.
	maxStop := config.FlightStopMaxNumber

	findFlightsDFS(searchData.OriginCityID, searchData.DestinationCityID, maxStop, make([]*entity.Flight, 0), make([]string, 0))

	if len(flightsMap) == 0 {
		return FlightVariantsResponse(c, flightsMap, err, 404, "Flights not found")
	}
	return FlightVariantsResponse(c, flightsMap, err, 201, "Flights not found")
}

func findFlightsDFS(originCityID string, destinationCityID string, stops int, flights []*entity.Flight, citiesID []string) {

	contains := slices.Contains(citiesID, originCityID) // Проверка на то что город уже есть в списке. Чтобы избежать цикличных вариантов перелета. Например: Алматы -> Астана -> Алматы

	if originCityID == destinationCityID {
		flightsMap = append(flightsMap, flights)
		return
	} else if stops <= 0 || contains == true {
		flights = nil
		return
	}

	nextFlights, err := repository.GetFlightsByOriginCity(originCityID)
	if err != nil {
		return
	}

	for i := 0; i < len(nextFlights); i++ {
		currentFlight := nextFlights[i]
		flights = append(flights, currentFlight)
		citiesID = append(citiesID, originCityID)
		findFlightsDFS(currentFlight.Direction.DestinationCity.ID, destinationCityID, stops-1, flights, citiesID)
		flights = flights[:len(flights)-1]
	}
	return
}
