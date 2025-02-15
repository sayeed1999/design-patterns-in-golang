package datasourcedecorators

import (
	"fmt"

	"github.com/sayeed1999/design-patterns-in-golang/decorator-pattern/datasource"
)

// Encryption Decorator adds encryption behavior
type EncryptionDecorator struct {
	*DataSourceDecorator // **Composition
}

func NewEncryptionDecorator(source datasource.DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{NewDataSourceDecorator(source)}
}

// Modifying behavior
func (e *EncryptionDecorator) WriteData(data string) {
	fmt.Println("Encrypting data...")
	e.wrappee.WriteData("[Encrypted] " + data)
}

// Modifying behavior
func (e *EncryptionDecorator) ReadData() string {
	fmt.Println("Decrypting data...")
	return e.wrappee.ReadData()
}
