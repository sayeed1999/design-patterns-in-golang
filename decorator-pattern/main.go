package main

import (
	"github.com/sayeed1999/design-patterns-in-golang/decorator-pattern/datasource"
	datasourcedecorators "github.com/sayeed1999/design-patterns-in-golang/decorator-pattern/datasource-decorators"
)

// Application demonstrates decorator usage.
func main() {
	// Simple usage without decorator
	fileDataSource := datasource.NewFileDataSource("somefile.dat")
	fileDataSource.WriteData("salaryRecords")

	// Adding data compression
	compressedSource := datasourcedecorators.NewCompressionDecorator(fileDataSource)
	compressedSource.WriteData("salaryRecords")

	/// Note:- Here you can pass the decorator instead of source because both implements the same interface!
	// Adding data encryption on top of compression
	encryptedSource := datasourcedecorators.NewEncryptionDecorator(compressedSource)
	encryptedSource.WriteData("salaryRecords")

	/// See how smoothly it atfirst decrypts data, then decompresses data!!!
	encryptedSource.ReadData()
}
