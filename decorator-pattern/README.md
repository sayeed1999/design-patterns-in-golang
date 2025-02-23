# Decorator Pattern

The Decorator Pattern is a structural design pattern that allows behavior to be added to individual objects, either statically or dynamically, without affecting the behavior of other objects from the same class. This pattern is useful for extending the functionalities of classes in a flexible and reusable way.

## Problem

Imagine you have a data source that needs to support additional functionalities like compression and encryption. Without a flexible design, you would need to create multiple classes for each combination of functionalities, leading to a complex and unmaintainable codebase.

## Solution

The Decorator Pattern helps solve this problem by allowing you to dynamically add functionalities to objects. In this example, decorators are used to add compression and encryption functionalities to a data source.

## Approach

1. **Define the Component Interface**: This interface will have methods that both the base component and decorators will implement.
2. **Create Concrete Component**: Implement the `Component` interface for the base component (e.g., FileDataSource).
3. **Create Decorator Class**: Implement the `Component` interface for the decorator class, which will wrap the base component and add additional functionalities.
4. **Create Concrete Decorators**: Implement the `Decorator` class for specific functionalities (e.g., CompressionDecorator, EncryptionDecorator).

### Structure

- **Component Interface**: Defines the methods that both the base component and decorators will implement.
  - [datasource/datasource.go](datasource/datasource.go)
- **Concrete Component**: FileDataSource class implementing the `Component` interface.
  - [datasource/file_data_source.go](datasource/file_data_source.go)
- **Decorator Class**: Base decorator class implementing the `Component` interface and wrapping the base component.
  - [datasource-decorators/decorator.go](datasource-decorators/decorator.go)
- **Concrete Decorators**: CompressionDecorator and EncryptionDecorator classes implementing the `Decorator` class.
  - [datasource-decorators/compression_decorator.go](datasource-decorators/compression_decorator.go)
  - [datasource-decorators/encryption_decorator.go](datasource-decorators/encryption_decorator.go)

### Usage

The main function demonstrates how to use the `FileDataSource` with decorators to add compression and encryption functionalities.

## Similarity with Proxy Pattern

Both Proxy & Decorator might seem similar, but Proxy can only forward requests.
But decorator might add or remove behavior.

E.g if there are 3 methods in the main class, proxy class must also implement the exact 03.
But a decorator class can contain 5 or 6 methods or even 2 based on need.

## Special Thanks

Tutorial Link: <https://refactoring.guru/design-patterns/decorator>
