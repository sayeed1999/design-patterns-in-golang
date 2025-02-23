package computer

import "github.com/sayeed1999/design-patterns-in-golang/bridge-pattern/printer"

type Computer interface {
	SetPrinter(printer.Printer)
	Print()
}
