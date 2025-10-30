package week5

import "fmt"

// FileReader - gets data from files
type FileReader struct{}

func (f *FileReader) read(filename string) string {
	fmt.Printf("1. FileReader: Opening file '%s'\n", filename)
	return "file data"
}

// Compressor - compressing data
type Compressor struct{}

func (c *Compressor) compress(data string) {
	fmt.Printf("2. Compressor: Compressing data: %s\n", data)
}

// Encryptor - encoding data
type Encryptor struct{}

func (e *Encryptor) encrypt(data string) {
	fmt.Printf("3. Encryptor: Encoding data: %s\n", data)
}
