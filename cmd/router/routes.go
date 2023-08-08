package router

import (
	"aviatoV3/internal/services/handlers"
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
	airlines.Get("/", handlers.GetAllAirlines)
	airlines.Get("/:id", handlers.GetSingleAirline)
	airlines.Post("/", handlers.CreateAirline)
	airlines.Put("/:id", handlers.UpdateAirline)
	airlines.Delete("/:id", handlers.DeleteAirline)

	// booking
	bookings.Get("/", handlers.GetAllBookings)
	bookings.Get("/:id", handlers.GetSingleBooking)
	bookings.Post("/", handlers.CreateBooking)
	bookings.Put("/:id", handlers.UpdateBooking)
	bookings.Delete("/:id", handlers.DeleteBooking)

	// city
	cities.Get("/", handlers.GetAllCities)
	cities.Get("/:id", handlers.GetSingleCity)
	cities.Post("/", handlers.CreateCity)
	cities.Put("/:id", handlers.UpdateCity)
	cities.Delete("/:id", handlers.DeleteCity)

	// country
	countries.Get("/", handlers.GetAllCountries)
	countries.Get("/:id", handlers.GetSingleCountry)
	countries.Post("/", handlers.CreateCountry)
	countries.Put("/:id", handlers.UpdateCountry)
	countries.Delete("/:id", handlers.DeleteCountry)

	// direction
	directions.Get("/", handlers.GetAllDirections)
	directions.Get("/:id", handlers.GetSingleDirection)
	directions.Post("/", handlers.CreateDirection)
	directions.Put("/:id", handlers.UpdateDirection)
	directions.Delete("/:id", handlers.DeleteDirection)

	// flight
	flights.Get("/", handlers.GetAllFlights)
	flights.Get("/:id", handlers.GetSingleFlight)
	flights.Post("/", handlers.CreateFlight)
	flights.Put("/:id", handlers.UpdateFlight)
	flights.Delete("/:id", handlers.DeleteFlight)

	// passenger
	passengers.Get("/", handlers.GetAllPassengers)
	passengers.Get("/:id", handlers.GetSinglePassenger)
	passengers.Post("/", handlers.CreatePassenger)
	passengers.Put("/:id", handlers.UpdatePassenger)
	passengers.Delete("/:id", handlers.DeletePassenger)

	// main
	/*
		api.Get("/", handlers.GetFlightsByOriginAndDestination)

	*/

}
