# Bridge Pattern

The Bridge Pattern is a structural design pattern that decouples an abstraction from its implementation, allowing the two to vary independently. This pattern is useful when both the abstraction and the implementation may have multiple variations and need to be extended independently.

**Note**: <i>At the end of this README.md file, I referenced the tutorial that I followed</i>

## Key Concepts

- **Abstraction**: Defines the interface for the control part of the two-class hierarchy. It maintains a reference to an object of type `Implementor`.
- **Implementor**: Defines the interface for the implementation part of the two-class hierarchy. This interface doesn't need to correspond directly to the `Abstraction` interface and can be very different.
- **Refined Abstraction**: Extends the interface defined by `Abstraction`.
- **Concrete Implementor**: Implements the `Implementor` interface and defines its concrete implementation.

## Example in This Repository

In this example, we demonstrate the Bridge Pattern using a scenario where different types of computers (Mac and Windows) can use different types of printers (HP and Epson).

### Structure

- **Printer Interface**: Defines the `PrintFile` method.
  - [printer/printer.go](printer/printer.go)
- **Concrete Implementors**: HP and Epson printers implementing the `Printer` interface.
  - [printer/hp.go](printer/hp.go)
  - [printer/epson.go](printer/epson.go)
- **Computer Interface**: Defines methods to set a printer and print.
  - [computer/computer.go](computer/computer.go)
- **Refined Abstractions**: Mac and Windows computers implementing the `Computer` interface.
  - [computer/mac.go](computer/mac.go)
  - [computer/windows.go](computer/windows.go)

### Usage

The main function demonstrates how different computers can use different printers interchangeably.

- [main.go](main.go)

```go
func main() {
    epsonPrinter := &printer.EpsonPrinter{}
    hpPrinter := &printer.HpPrinter{}

    mac := &computer.Mac{}

    mac.SetPrinter(epsonPrinter)
    mac.Print()

    mac.SetPrinter(hpPrinter)
    mac.Print()

    windows := &computer.Windows{}

    windows.SetPrinter(epsonPrinter)
    windows.Print()

    windows.SetPrinter(hpPrinter)
    windows.Print()
}
```

### Output

```bash
Printing file with a EPSON printer...
Printing file with a HP printer...
Printing file with a EPSON printer...
Printing file with a HP printer...
```

## Benefits

- **Decoupling Abstraction and Implementation**: Allows both to vary independently.
- **Increased Flexibility**: You can change or extend the abstraction and implementation hierarchies independently.
- **Simplified Code Maintenance**: Changes in the implementation do not affect the client code.

## Special Thanks

- [Refactoring Guru: Bridge Pattern](https://refactoring.guru/design-patterns/bridge)

The example problem and solution at Refactoring Guru of RedCircle, RedSquare, BlueCircle, BlueSquare I found so great to learn! I recommend readers to check their tutorial for better learning!
