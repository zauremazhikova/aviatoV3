package handlers

import (
	"aviatoV3/internal/entities"
	"aviatoV3/internal/repositories"
	"aviatoV3/internal/services/transport"
	"github.com/gofiber/fiber/v2"
)

func GetAllBookings(c *fiber.Ctx) error {

	responseBookings, err := repositories.GetBookings()
	return transport.ResponseBookings(c, responseBookings, err)

}

func GetSingleBooking(c *fiber.Ctx) error {

	id := c.Params("id")
	booking, err := repositories.GetBooking(id)
	return transport.ResponseBooking(c, booking, err)

}

func CreateBooking(c *fiber.Ctx) error {

	insertStruct, err := transport.ValidateBookingInsertData(c)
	if err != nil {
		return err
	}

	passenger, err := repositories.GetPassenger(insertStruct.PassengerID)
	err = transport.ResponseBookingPassengerNotFound(c, passenger, err)
	if err != nil {
		return err
	}

	flight, err := repositories.GetFlight(insertStruct.FlightID)
	err = transport.ResponseBookingFlightNotFound(c, flight, err)
	if err != nil {
		return err
	}

	booking := new(entities.Booking)
	booking.BookingNumber = insertStruct.BookingNumber
	booking.Passenger = *passenger
	booking.Flight = *flight

	err = repositories.CreateBooking(booking)
	return transport.ResponseBookingCreate(c, err)

}

func UpdateBooking(c *fiber.Ctx) error {

	id := c.Params("id")
	booking, err := repositories.GetBooking(id)

	err = transport.ResponseBookingNotFound(c, booking, err)
	if err != nil {
		return err
	}

	updateStruct, err := transport.ValidateBookingUpdateData(c)
	if err != nil {
		return err
	}

	passenger, err := repositories.GetPassenger(updateStruct.PassengerID)
	err = transport.ResponseBookingPassengerNotFound(c, passenger, err)
	if err != nil {
		return err
	}

	flight, err := repositories.GetFlight(updateStruct.FlightID)
	err = transport.ResponseBookingFlightNotFound(c, flight, err)
	if err != nil {
		return err
	}

	booking.BookingNumber = updateStruct.BookingNumber
	booking.Passenger = *passenger
	booking.Flight = *flight

	err = repositories.UpdateBooking(booking)

	return transport.ResponseBookingUpdate(c, err)
}

func DeleteBooking(c *fiber.Ctx) error {

	id := c.Params("id")
	booking, err := repositories.GetBooking(id)

	err = transport.ResponseBookingNotFound(c, booking, err)
	if err != nil {
		return err
	}

	err = repositories.DeleteBooking(id)
	return transport.ResponseBookingDelete(c, err)

}
