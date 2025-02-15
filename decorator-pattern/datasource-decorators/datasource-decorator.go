package datasourcedecorators

import "github.com/sayeed1999/design-patterns-in-golang/decorator-pattern/datasource"

/// Note: -
// Do I really need a base decorator??
// Normally the decorators have no code to share! E.g encryption & compression are totally different logic.
// Removing the base decorator can simplify the code & make easier to follow.

// DataSourceDecorator is the base decorator class.
type DataSourceDecorator struct {
	wrappee datasource.DataSource // **Aggregation
}

func NewDataSourceDecorator(source datasource.DataSource) *DataSourceDecorator {
	return &DataSourceDecorator{wrappee: source}
}

/// The base decorator class must implement the exact implementations of the main class

// Forwarding write data of base
func (d *DataSourceDecorator) WriteData(data string) {
	d.wrappee.WriteData(data)
}

// Forwarding read data of base
func (d *DataSourceDecorator) ReadData() string {
	return d.wrappee.ReadData()
}
