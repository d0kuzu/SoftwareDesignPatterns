package printers

import "fmt"

// Epson - concrete implementer of printer interface
type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
