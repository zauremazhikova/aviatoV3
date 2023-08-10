package entityHandlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transports/entityTransports"
	"github.com/gofiber/fiber/v2"
)

func GetAllBookings(c *fiber.Ctx) error {

	responseBookings, err := repositories.GetBookings()
	return entityTransports.ResponseBookings(c, responseBookings, err)

}

func GetSingleBooking(c *fiber.Ctx) error {

	id := c.Params("id")
	booking, err := repositories.GetBooking(id)
	return entityTransports.ResponseBooking(c, booking, err)

}

func CreateBooking(c *fiber.Ctx) error {

	insertStruct, err := entityTransports.ValidateBookingInsertData(c)
	if err != nil {
		return err
	}

	passenger, err := repositories.GetPassenger(insertStruct.PassengerID)
	err = entityTransports.ResponseBookingPassengerNotFound(c, passenger, err)
	if err != nil {
		return err
	}

	flight, err := repositories.GetFlight(insertStruct.FlightID)
	err = entityTransports.ResponseBookingFlightNotFound(c, flight, err)
	if err != nil {
		return err
	}

	booking := new(entities.Booking)
	booking.BookingNumber = insertStruct.BookingNumber
	booking.Passenger = *passenger
	booking.Flight = *flight

	err = repositories.CreateBooking(booking)
	return entityTransports.ResponseBookingCreate(c, err)

}

func UpdateBooking(c *fiber.Ctx) error {

	id := c.Params("id")
	booking, err := repositories.GetBooking(id)

	err = entityTransports.ResponseBookingNotFound(c, booking, err)
	if err != nil {
		return err
	}

	updateStruct, err := entityTransports.ValidateBookingUpdateData(c)
	if err != nil {
		return err
	}

	passenger, err := repositories.GetPassenger(updateStruct.PassengerID)
	err = entityTransports.ResponseBookingPassengerNotFound(c, passenger, err)
	if err != nil {
		return err
	}

	flight, err := repositories.GetFlight(updateStruct.FlightID)
	err = entityTransports.ResponseBookingFlightNotFound(c, flight, err)
	if err != nil {
		return err
	}

	booking.BookingNumber = updateStruct.BookingNumber
	booking.Passenger = *passenger
	booking.Flight = *flight

	err = repositories.UpdateBooking(booking)

	return entityTransports.ResponseBookingUpdate(c, err)
}

func DeleteBooking(c *fiber.Ctx) error {

	id := c.Params("id")
	booking, err := repositories.GetBooking(id)

	err = entityTransports.ResponseBookingNotFound(c, booking, err)
	if err != nil {
		return err
	}

	err = repositories.DeleteBooking(id)
	return entityTransports.ResponseBookingDelete(c, err)

}
