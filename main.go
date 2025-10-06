package main

import (
	. "SoftwareDesignPatterns/week4/abstractions/computers"
	. "SoftwareDesignPatterns/week4/implementers/printers"
	"fmt"
)

func main() {
	// Creating implementers (printers)
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	// Creating abstractions (computers)
	macComputer := &Mac{}
	winComputer := &Windows{}

	// 1. Using Mac computer with Hp printer
	fmt.Println("--- Mac with HP ---")
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	// 2. Using Mac computer with Epson printer
	fmt.Println("--- Mac with Epson ---")
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	// 3. Using Windows computer with Hp printer
	fmt.Println("--- Windows with HP ---")
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	// 4. Using Windows computer with Epson printer
	fmt.Println("--- Windows with Epson ---")
	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()

}
