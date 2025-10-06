package printers

import "fmt"

// Hp - concrete implementer of printer interface
type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
