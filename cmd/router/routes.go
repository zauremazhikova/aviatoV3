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
	//bookings := api.Group("/booking")
	cities := api.Group("/city")
	countries := api.Group("/country")
	directions := api.Group("/direction")
	//flights := api.Group("/flight")
	passengers := api.Group("/passenger")

	// airline
	airlines.Get("/", service.GetAllAirlines)
	airlines.Get("/:id", service.GetSingleAirline)
	airlines.Post("/", service.CreateAirline)
	airlines.Put("/:id", service.UpdateAirline)
	airlines.Delete("/:id", service.DeleteAirline)

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

	// passenger
	passengers.Get("/", service.GetAllPassengers)
	passengers.Get("/:id", service.GetSinglePassenger)
	passengers.Post("/", service.CreatePassenger)
	passengers.Put("/:id", service.UpdatePassenger)
	passengers.Delete("/:id", service.DeletePassenger)

	/*
		// booking
		bookings.Get("/", entityHandlers.GetAllBookings)
		bookings.Get("/:id", entityHandlers.GetSingleBooking)
		bookings.Post("/", entityHandlers.CreateBooking)
		bookings.Put("/:id", entityHandlers.UpdateBooking)
		bookings.Delete("/:id", entityHandlers.DeleteBooking)

		// flight
		flights.Get("/", entityHandlers.GetAllFlights)
		flights.Get("/:id", entityHandlers.GetSingleFlight)
		flights.Post("/", entityHandlers.CreateFlight)
		flights.Put("/:id", entityHandlers.UpdateFlight)
		flights.Delete("/:id", entityHandlers.DeleteFlight)

	*/

	// main
	/*
		api.Get("/", handlers.GetFlightsByOriginAndDestination)

	*/

}
