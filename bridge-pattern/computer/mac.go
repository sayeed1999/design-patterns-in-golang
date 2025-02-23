package computer

import "github.com/sayeed1999/design-patterns-in-golang/bridge-pattern/printer"

// implements computer
type Mac struct {
	printer printer.Printer
}

// method injection over ctor injection!
func (m *Mac) SetPrinter(printer printer.Printer) {
	m.printer = printer
}

func (m *Mac) Print() {
	m.printer.PrintFile()
}
