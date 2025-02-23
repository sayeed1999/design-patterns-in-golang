# Bridge Pattern

The Bridge Pattern is a structural design pattern that decouples an abstraction from its implementation, allowing the two to vary independently. This pattern is useful when both the abstraction and the implementation may have multiple variations and need to be extended independently.

**Note**: <i>At the end of this README.md file, I referenced the tutorial that I followed</i>

## Problem

Imagine you have different types of computers (e.g., Mac and Windows) that need to print documents using different types of printers (e.g., HP and Epson). You want to be able to mix and match computers and printers without creating a complex and tightly coupled codebase.

## Solution

The Bridge Pattern helps solve this problem by decoupling the abstraction (Computer) from its implementation (Printer). This allows you to extend both independently without affecting each other.

## Approach

1. **Define the Printer Interface**: This interface will have a method `PrintFile` that all concrete printers will implement.
2. **Create Concrete Printers**: Implement the `Printer` interface for different printer types (e.g., HP and Epson).
3. **Define the Computer Interface**: This interface will have methods to set a printer and print a document.
4. **Create Concrete Computers**: Implement the `Computer` interface for different computer types (e.g., Mac and Windows).
5. **Bridge the Abstraction and Implementation**: Use the `Computer` interface to set different `Printer` implementations and print documents.

### Structure

- **Printer Interface**: Defines the `PrintFile` method.
  - [printer/printer.go](printer/printer.go)
- **Concrete Printers**: HP and Epson printers implementing the `Printer` interface.
  - [printer/hp.go](printer/hp.go)
  - [printer/epson.go](printer/epson.go)
- **Computer Interface**: Defines methods to set a printer and print.
  - [computer/computer.go](computer/computer.go)
- **Concrete Computers**: Mac and Windows computers implementing the `Computer` interface.
  - [computer/mac.go](computer/mac.go)
  - [computer/windows.go](computer/windows.go)

### Usage

The main function demonstrates how different computers can use different printers interchangeably.

- [main.go](main.go)

### Output

```bash
Printing file with a EPSON printer...
Printing file with a HP printer...
Printing file with a EPSON printer...
Printing file with a HP printer...
```

## Benefits

- Decoupling Abstraction and Implementation: Allows both to vary independently.
- Increased Flexibility: You can change or extend the abstraction and implementation hierarchies independently.
- Simplified Code Maintenance: Changes in the implementation do not affect the client code.

## Special Thanks

[Refactoring Guru: Bridge Pattern](https://refactoring.guru/design-patterns/bridge)

The example problem and solution at Refactoring Guru of `RedCircle`, `RedSquare`, `BlueCircle`, `BlueSquare` I found so great to learn! I recommend readers to check their tutorial for better learning!
