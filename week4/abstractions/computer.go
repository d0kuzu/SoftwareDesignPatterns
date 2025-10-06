package abstractions

import "SoftwareDesignPatterns/week4/implementers"

type Computer interface {
	Print()
	SetPrinter(p implementers.Printer)
}
