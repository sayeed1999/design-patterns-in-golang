package computer

import "github.com/sayeed1999/design-patterns-in-golang/bridge-pattern/printer"

// implements computer
type Windows struct {
	printer printer.Printer
}

// method injection over ctor injection!
func (m *Windows) SetPrinter(printer printer.Printer) {
	m.printer = printer
}

func (m *Windows) Print() {
	m.printer.PrintFile()
}
