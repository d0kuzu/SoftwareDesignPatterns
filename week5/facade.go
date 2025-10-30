package week5

import "fmt"

// FileProcessorFacade - facade that simplifies interface
type FileProcessorFacade struct {
	reader     *FileReader
	compressor *Compressor
	encryptor  *Encryptor
}

// NewFileProcessorFacade - facade constructor
func NewFileProcessorFacade() *FileProcessorFacade {
	return &FileProcessorFacade{
		reader:     &FileReader{},
		compressor: &Compressor{},
		encryptor:  &Encryptor{},
	}
}

// ProcessFile - facade's method that simplified using of subsystem
func (f *FileProcessorFacade) ProcessFile(filename string, method MethodType) {
	fmt.Println("--- Starting facade high-level process: ProcessFile ---")

	// 1. Reading file
	fileData := f.reader.read(filename)

	// 2. Understand the type of action needed
	if method == Compress {
		f.compressor.compress(fileData)
	} else if method == Encrypt {
		f.encryptor.encrypt(fileData)
	} else {
		fmt.Println("   ... Method not recognized.")
	}

	fmt.Println("--- Operation ProcessFile ended ---")
}
