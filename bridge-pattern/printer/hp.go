package printer

import "fmt"

// implements printer
type HpPrinter struct{}

func (p *HpPrinter) PrintFile() {
	fmt.Println("Printing file with a HP printer...")
}
