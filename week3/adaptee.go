package week3

import "strconv"

// StringService - adaptee structure, that work only with string input and returns string
type StringService struct{}

// AddStrings выполняет сложение, но оперирует только строками.
func (s *StringService) AddStrings(strA string, strB string) string {
	if strA == "" || strB == "" {
		return "error: missing operand"
	}

	// Converting to int
	valA, errA := strconv.Atoi(strA)
	valB, errB := strconv.Atoi(strB)

	if errA != nil || errB != nil {
		return "error: invalid input"
	}

	result := valA + valB       // sum
	return strconv.Itoa(result) // converting and returning sum as string
}
