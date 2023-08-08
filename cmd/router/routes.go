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
	//bookings := api.Group("/booking")
	cities := api.Group("/city")
	countries := api.Group("/country")
	//directions := api.Group("/direction")
	//flights := api.Group("/flight")
	passengers := api.Group("/passenger")

	// airline
	airlines.Get("/", handlers.GetAllAirlines)
	airlines.Get("/:id", handlers.GetSingleAirline)
	airlines.Post("/", handlers.CreateAirline)
	airlines.Put("/:id", handlers.UpdateAirline)
	airlines.Delete("/:id", handlers.DeleteAirline)

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

	// passenger
	passengers.Get("/", handlers.GetAllPassengers)
	passengers.Get("/:id", handlers.GetSinglePassenger)
	passengers.Post("/", handlers.CreatePassenger)
	passengers.Put("/:id", handlers.UpdatePassenger)
	passengers.Delete("/:id", handlers.DeletePassenger)

	// main
	/*
		api.Get("/", handlers.GetFlightsByOriginAndDestination)

		// booking
		bookings.Get("/", booking.GetAll)
		bookings.Get("/:id", booking.GetSingle)
		bookings.Post("/", booking.Create)
		bookings.Put("/:id", booking.Update)
		bookings.Delete("/:id", booking.Delete)

		// direction
		directions.Get("/", direction.GetAll)
		directions.Get("/:id", direction.GetSingle)
		directions.Post("/", direction.Create)
		directions.Put("/:id", direction.Update)
		directions.Delete("/:id", direction.Delete)

		// flight
		flights.Get("/", flight.GetAll)
		flights.Get("/:id", flight.GetSingle)
		flights.Post("/", flight.Create)
		flights.Put("/:id", flight.Update)
		flights.Delete("/:id", flight.Delete)

	*/

}
