package datasource

import "fmt"

// Implements DataSource.
// FileDataSource is a concrete component providing default implementations.
type FileDataSource struct {
	filename string
}

func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{filename: filename}
}

func (f *FileDataSource) WriteData(data string) {
	fmt.Printf("Writing data to file: %s\n", data)
}

func (f *FileDataSource) ReadData() string {
	fmt.Println("Reading data from file")
	return "file data"
}
