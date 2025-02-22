package main

import "fmt"

type ShippingService struct {
	order *Order
}

func (s *ShippingService) AcceptOrder(order *Order) {
	s.order = order
}

func (s *ShippingService) CalculateShippingExpenses() {
	if s.order != nil {
		s.order.ShippingPrice = 15.5
	}
}

func (s *ShippingService) ShipOrder() {
	if s.order != nil {
		fmt.Println(s.order.ToString())
		fmt.Printf("Order is being shipped to %s...\n", s.order.ShippingAddress)
	} else {
		fmt.Println("No order to ship.")
	}
}
