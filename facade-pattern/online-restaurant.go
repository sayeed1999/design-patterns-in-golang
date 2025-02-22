package main

import "fmt"

type OnlineRestaurant struct {
	cart []*Order
}

func NewOnlineRestaurant() *OnlineRestaurant {
	return &OnlineRestaurant{
		cart: make([]*Order, 0),
	}
}

func (or *OnlineRestaurant) AddOrderToCart(order *Order) {
	or.cart = append(or.cart, order)
}

func (or *OnlineRestaurant) CompleteOrders() {
	fmt.Println("Orders completed. Dispatch in progress...")
}
