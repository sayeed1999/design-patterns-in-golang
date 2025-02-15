package datasourcedecorators

import (
	"fmt"

	"github.com/sayeed1999/design-patterns-in-golang/decorator-pattern/datasource"
)

type CompressionDecorator struct {
	*DataSourceDecorator // **Composition
}

func NewCompressionDecorator(source datasource.DataSource) *CompressionDecorator {
	dataSourceDecorator := NewDataSourceDecorator(source)
	return &CompressionDecorator{dataSourceDecorator}
}

// Modifying behavior
func (c *CompressionDecorator) WriteData(data string) {
	fmt.Println("Compressing data...")
	compressed := "[Compressed] " + data
	c.wrappee.WriteData(compressed)
}

// Modifying behavior
func (d *CompressionDecorator) ReadData() string {
	fmt.Println("Decompressing data...")
	return d.wrappee.ReadData()
}
