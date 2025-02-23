# Factory Pattern

The Factory Pattern is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. This pattern is useful for creating objects without specifying the exact class of object that will be created.

## Problem

Imagine you have a payment processing system that supports multiple payment methods (e.g., Debit Card, PayPal, Bank Transfer). You want to add support for new payment methods (e.g., Crypto) without modifying the existing codebase. Without a flexible design, adding new payment methods would require changes to the existing code, making it complex and difficult to maintain.

## Solution

The Factory Pattern helps solve this problem by defining a factory method that creates objects based on input parameters. In this example, a `PaymentFactory` function is created to return the appropriate payment processor based on the payment type.

## Approach

1. **Define the Product Interface**: This interface will have methods that all concrete products (payment processors) will implement.
2. **Create Concrete Products**: Implement the `Product` interface for different payment methods (e.g., DebitCard, PayPal, BankTransfer).
3. **Create Factory Method**: Implement a factory method that returns the appropriate product based on input parameters.

### Structure

- **Product Interface**: Defines the methods that all concrete products will implement.
  - [payment/payment.go](payment/payment.go)
- **Concrete Products**: Implement the `Product` interface for different payment methods.
  - [payment/debit_card.go](payment/debit_card.go)
  - [payment/paypal.go](payment/paypal.go)
  - [payment/bank_transfer.go](payment/bank_transfer.go)
  - [payment/crypto.go](payment/crypto.go) (to be implemented)
- **Factory Method**: Returns the appropriate product based on input parameters.
  - [payment/factory.go](payment/factory.go)

### Usage

The main function demonstrates how to use the `PaymentFactory` to process payments using different payment methods.

- [main.go](main.go)

## Benefits

- **Flexibility**: Allows you to add new products (payment methods) without modifying the existing code.
- **Simplified Code Maintenance**: Keeps the client code clean and focused on its primary responsibilities.
- **Scalability**: Easily extendable to include new products without changing the client code.

## Refenrences

Read my article on Medium on Factory Pattern: -

- [Using the Factory Design Pattern in Post Upvoting System for Stack Overflow📚](https://medium.com/@sayeedrahman_67698/using-the-factory-design-pattern-in-post-upvoting-system-for-stack-overflow-c21afee56c94)
- [Factory Design Pattern in Go: Simplifying Payment Gateways](https://medium.com/@sayeedrahman_67698/factory-design-pattern-in-go-simplifying-payment-gateways-ce2e65e78433)
