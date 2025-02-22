package main

import "fmt"

func main() {

	restaurant := NewOnlineRestaurant()
	shippingService := &ShippingService{}

	chickenOrder := &Order{
		DishName:        "Chicken with rice",
		DishPrice:       20.0,
		User:            "User1",
		ShippingAddress: "Random street 123",
	}

	sushiOrder := &Order{
		DishName:        "Sushi",
		DishPrice:       52.0,
		User:            "User2",
		ShippingAddress: "More random street 321"}

	// ------ Without facade, this is how the main func looks ------

	restaurant.AddOrderToCart(chickenOrder)
	restaurant.AddOrderToCart(sushiOrder)
	restaurant.CompleteOrders()
	fmt.Println()

	shippingService.AcceptOrder(chickenOrder)
	shippingService.CalculateShippingExpenses()
	shippingService.ShipOrder()
	fmt.Println()

	shippingService.AcceptOrder(sushiOrder)
	shippingService.CalculateShippingExpenses()
	shippingService.ShipOrder()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// With facade, it should be oneliner & all internal logic abstracted away!
	facade := &Facade{
		restaurant:      restaurant,
		shippingService: shippingService,
	}

	orders := []*Order{
		chickenOrder,
		sushiOrder,
	}

	facade.OrderFood(orders)
}
