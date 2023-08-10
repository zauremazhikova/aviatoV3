package router

import (
	"aviatoV3/internal/services/transports/entityTransports"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	airlines := api.Group("/airline")
	//bookings := api.Group("/booking")
	//cities := api.Group("/city")
	//countries := api.Group("/country")
	//directions := api.Group("/direction")
	//flights := api.Group("/flight")
	//passengers := api.Group("/passenger")

	// airline
	airlines.Get("/", entityTransports.GetAllAirlines)

	airlines.Get("/:id", entityTransports.GetSingleAirline)
	airlines.Post("/", entityTransports.CreateAirline)
	airlines.Put("/:id", entityTransports.UpdateAirline)
	airlines.Delete("/:id", entityTransports.DeleteAirline)
	/*
		// booking
		bookings.Get("/", entityHandlers.GetAllBookings)
		bookings.Get("/:id", entityHandlers.GetSingleBooking)
		bookings.Post("/", entityHandlers.CreateBooking)
		bookings.Put("/:id", entityHandlers.UpdateBooking)
		bookings.Delete("/:id", entityHandlers.DeleteBooking)

		// city
		cities.Get("/", entityHandlers.GetAllCities)
		cities.Get("/:id", entityHandlers.GetSingleCity)
		cities.Post("/", entityHandlers.CreateCity)
		cities.Put("/:id", entityHandlers.UpdateCity)
		cities.Delete("/:id", entityHandlers.DeleteCity)

		// country
		countries.Get("/", entityHandlers.GetAllCountries)
		countries.Get("/:id", entityHandlers.GetSingleCountry)
		countries.Post("/", entityHandlers.CreateCountry)
		countries.Put("/:id", entityHandlers.UpdateCountry)
		countries.Delete("/:id", entityHandlers.DeleteCountry)

		// direction
		directions.Get("/", entityHandlers.GetAllDirections)
		directions.Get("/:id", entityHandlers.GetSingleDirection)
		directions.Post("/", entityHandlers.CreateDirection)
		directions.Put("/:id", entityHandlers.UpdateDirection)
		directions.Delete("/:id", entityHandlers.DeleteDirection)

		// flight
		flights.Get("/", entityHandlers.GetAllFlights)
		flights.Get("/:id", entityHandlers.GetSingleFlight)
		flights.Post("/", entityHandlers.CreateFlight)
		flights.Put("/:id", entityHandlers.UpdateFlight)
		flights.Delete("/:id", entityHandlers.DeleteFlight)

		// passenger
		passengers.Get("/", entityHandlers.GetAllPassengers)
		passengers.Get("/:id", entityHandlers.GetSinglePassenger)
		passengers.Post("/", entityHandlers.CreatePassenger)
		passengers.Put("/:id", entityHandlers.UpdatePassenger)
		passengers.Delete("/:id", entityHandlers.DeletePassenger)*/

	// main
	/*
		api.Get("/", handlers.GetFlightsByOriginAndDestination)

	*/

}
