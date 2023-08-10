package router

import (
	"aviatoV3/internal/service"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	airlines := api.Group("/airline")
	bookings := api.Group("/booking")
	cities := api.Group("/city")
	countries := api.Group("/country")
	directions := api.Group("/direction")
	flights := api.Group("/flight")
	passengers := api.Group("/passenger")

	// airline
	airlines.Get("/", service.GetAllAirlines)
	airlines.Get("/:id", service.GetSingleAirline)
	airlines.Post("/", service.CreateAirline)
	airlines.Put("/:id", service.UpdateAirline)
	airlines.Delete("/:id", service.DeleteAirline)

	// booking
	bookings.Get("/", service.GetAllBookings)
	bookings.Get("/:id", service.GetSingleBooking)
	bookings.Post("/", service.CreateBooking)
	bookings.Put("/:id", service.UpdateBooking)
	bookings.Delete("/:id", service.DeleteBooking)

	// city
	cities.Get("/", service.GetAllCities)
	cities.Get("/:id", service.GetSingleCity)
	cities.Post("/", service.CreateCity)
	cities.Put("/:id", service.UpdateCity)
	cities.Delete("/:id", service.DeleteCity)

	// country
	countries.Get("/", service.GetAllCountries)
	countries.Get("/:id", service.GetSingleCountry)
	countries.Post("/", service.CreateCountry)
	countries.Put("/:id", service.UpdateCountry)
	countries.Delete("/:id", service.DeleteCountry)

	// direction
	directions.Get("/", service.GetAllDirections)
	directions.Get("/:id", service.GetSingleDirection)
	directions.Post("/", service.CreateDirection)
	directions.Put("/:id", service.UpdateDirection)
	directions.Delete("/:id", service.DeleteDirection)

	// flight
	flights.Get("/", service.GetAllFlights)
	flights.Get("/:id", service.GetSingleFlight)
	flights.Post("/", service.CreateFlight)
	flights.Put("/:id", service.UpdateFlight)
	flights.Delete("/:id", service.DeleteFlight)

	// passenger
	passengers.Get("/", service.GetAllPassengers)
	passengers.Get("/:id", service.GetSinglePassenger)
	passengers.Post("/", service.CreatePassenger)
	passengers.Put("/:id", service.UpdatePassenger)
	passengers.Delete("/:id", service.DeletePassenger)

	// main
	/*
		api.Get("/", handlers.GetFlightsByOriginAndDestination)

	*/

}
