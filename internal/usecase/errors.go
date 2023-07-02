package usecase

import "fmt"

type FlightNotFound struct {
	flightID int
}

func (f FlightNotFound) Error() string {
	return fmt.Sprint("Flight not found for flight id: ", f.flightID)
}

type FlightCapacityError struct {
	Capacity int
	Required int
}

func (f FlightCapacityError) Error() string {
	return fmt.Sprint("There are not enough capacity! Capcity: ", f.Capacity, " Required: ",
		f.Required)
}

type SomePassengerNotFound struct {
	TotalPassengers int
	FoundPassengers int
}

func (s SomePassengerNotFound) Error() string {
	return fmt.Sprint("Some passenger not found! TotalPassengers: ", s.TotalPassengers, " FoundPassengers: ",
		s.FoundPassengers)
}

type OrderNotFound struct {
	orderID int
}

func (o OrderNotFound) Error() string {
	return fmt.Sprint("Order not found for order id:", o.orderID)
}

type OrderNotPaid struct {
	orderID int
}

func (o OrderNotPaid) Error() string {
	return fmt.Sprint("Order [id: ", o.orderID, "] is not paid")
}

type OrdrAlreadyDelivered struct {
	orderID int
}

func (o OrdrAlreadyDelivered) Error() string {
	return fmt.Sprint("Order [id: ", o.orderID, "] already delivered")
}

type OrderItemsNotFound struct {
	orderID uint
}

func (o OrderItemsNotFound) Error() string {
	return fmt.Sprint("Order items not paid for order id:", o.orderID)

}

type InvalidBankName struct {
	name string
}

func (i InvalidBankName) Error() string {
	return fmt.Sprint("Invalid bank name: ", i.name)
}

type VerifyingPaymentFailed struct {
	orderID int
}

func (v VerifyingPaymentFailed) Error() string {
	return fmt.Sprint("Verifying payment failed for orderId: ", v.orderID)
}
