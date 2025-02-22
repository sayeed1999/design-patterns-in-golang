package main

import "fmt"

type Facade struct {
	restaurant      *OnlineRestaurant
	shippingService *ShippingService
}

func NewFacade(restaurant *OnlineRestaurant, shippingService *ShippingService) *Facade {
	return &Facade{
		restaurant:      restaurant,
		shippingService: shippingService,
	}
}

func (f *Facade) OrderFood(orders []*Order) {
	for _, order := range orders {
		f.restaurant.AddOrderToCart(order)
	}

	f.restaurant.CompleteOrders()
	fmt.Println()

	for _, order := range orders {
		f.shippingService.AcceptOrder(order)
		f.shippingService.CalculateShippingExpenses()
		f.shippingService.ShipOrder()
		fmt.Println()
	}
}
