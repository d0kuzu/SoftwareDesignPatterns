package computers

import (
	. "SoftwareDesignPatterns/week4/implementers"
	"fmt"
)

// Mac - Refined Abstraction
type Mac struct {
	printer Printer // Bridge-field
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile() // Delegation of printer's method
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}
