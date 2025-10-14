package main

import (
	. "SoftwareDesignPatterns/week4/abstractions/computers"
	. "SoftwareDesignPatterns/week4/implementers/printers"
	. "SoftwareDesignPatterns/week6"
)

func main() {
	// Creating facade
	processor := NewFileProcessorFacade()

	// Using his functionality without knowing subsystems complex structure
	processor.ProcessFile("report.docx", Compress)
	processor.ProcessFile("secret.txt", Encrypt)
}
