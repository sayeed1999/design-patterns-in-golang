package main

import "fmt"

type Order struct {
	DishName        string
	DishPrice       float64
	User            string
	ShippingAddress string
	ShippingPrice   float64
}

func (o *Order) ToString() string {
	return fmt.Sprintf("User %s ordered %s. The full price is %.2f dollars.",
		o.User, o.DishName, o.DishPrice+o.ShippingPrice)
}
