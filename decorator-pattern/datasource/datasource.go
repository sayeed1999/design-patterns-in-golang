package datasource

// DataSource defines base operations that can be altered by decorators.
type DataSource interface {
	WriteData(data string)
	ReadData() string
}
