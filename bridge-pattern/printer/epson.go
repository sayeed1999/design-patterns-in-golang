package printer

import "fmt"

type EpsonPrinter struct{}

func (p *EpsonPrinter) PrintFile() {
	fmt.Println("Printing file with a EPSON printer...")
}
