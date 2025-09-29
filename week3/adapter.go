package week3

import (
	"fmt"
	"strconv"
)

// Calculator - expected interface
type Calculator interface {
	Add(a int, b int) (int, error)
}

// ServiceAdapter - adapter which realises Calculator.
type ServiceAdapter struct {
	Service *StringService // adaptee
}

// Add realise expected method from Calculator.
func (adapter *ServiceAdapter) Add(a int, b int) (int, error) {
	// converting int to string
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)

	// adaptee method call
	resultStr := adapter.Service.AddStrings(strA, strB)

	// adapting output (string to int)
	if len(resultStr) > 6 && resultStr[:5] == "error" {
		return 0, fmt.Errorf("service error: %s", resultStr)
	}

	finalResult, err := strconv.Atoi(resultStr)
	if err != nil {
		return 0, fmt.Errorf("adapter failed to parse result: %w", err)
	}

	return finalResult, nil
}
