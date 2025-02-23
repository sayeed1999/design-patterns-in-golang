package main

import (
	"github.com/sayeed1999/design-patterns-in-golang/bridge-pattern/computer"
	"github.com/sayeed1999/design-patterns-in-golang/bridge-pattern/printer"
)

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
