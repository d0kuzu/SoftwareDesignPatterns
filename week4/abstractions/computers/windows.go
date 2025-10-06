package computers

import (
	. "SoftwareDesignPatterns/week4/implementers"
	"fmt"
)

// Windows - Refined Abstraction
type Windows struct {
	printer Printer // Bridge-field
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile() // Delegation of printer's method
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
