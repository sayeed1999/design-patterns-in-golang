
# Facade Pattern

The Facade Pattern is a structural design pattern that provides a simplified interface to a complex subsystem. It helps to hide the complexities of the system and provides an easy-to-use interface to the client.

## Problem

Imagine you have an online restaurant system where customers can place orders, and the orders need to be processed and shipped. Without a unified interface, the client code has to interact with multiple subsystems (e.g., restaurant and shipping service) directly, making the code complex and difficult to maintain.

## Solution

The Facade Pattern helps solve this problem by providing a unified interface to the complex subsystem. In this example, a `Facade` class is created to provide a simple interface for placing orders and handling the shipping process, hiding the complexities of interacting with each subsystem.

## Approach

1. **Identify the Subsystems**: Identify the various subsystems that need to be used together (e.g., restaurant and shipping service).
2. **Create the Facade**: Create a facade class that provides a simplified interface to the subsystems.
3. **Delegate Calls**: The facade class should delegate the calls to the appropriate subsystems to perform the task.

### Structure

- **Order Class**: Represents an order placed by a customer.
  - [order.go](order.go)
- **Restaurant Class**: Handles adding orders to the cart and completing orders.
  - [restaurant.go](restaurant.go)
- **Shipping Service Class**: Handles accepting orders, calculating shipping expenses, and shipping orders.
  - [shipping_service.go](shipping_service.go)
- **Facade Class**: Provides a simplified interface to the restaurant and shipping service.
  - [facade.go](facade.go)

### Usage

The main function demonstrates how to use the `Facade` to place orders and handle the shipping process.

- [main.go](main.go)

### Output

```bash
Adding order to cart: Chicken with rice
Adding order to cart: Sushi
Completing orders...
Order accepted: Chicken with rice
Calculating shipping expenses...
Shipping order: Chicken with rice
Order accepted: Sushi
Calculating shipping expenses...
Shipping order: Sushi

Adding order to cart: Chicken with rice
Adding order to cart: Sushi
Completing orders...
Order accepted: Chicken with rice
Calculating shipping expenses...
Shipping order: Chicken with rice
Order accepted: Sushi
Calculating shipping expenses...
Shipping order: Sushi
```

## Benefits

- Simplified Interface: Provides a - simplified interface to a complex subsystem.
- Reduced Complexity: Hides the complexities of the subsystem from the client.
- Improved Code Maintenance: Changes in the subsystem do not affect the client code.

## Special Thanks

Tutorial Link: <https://code-maze.com/facade/>

The tutorial at Code Maze provides a great example of the Facade Pattern. I recommend readers to check their tutorial for better learning!
