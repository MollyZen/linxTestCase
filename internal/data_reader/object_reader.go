package data_reader

type ObjectReader[K any] interface {
	// ReadNext Returns next value. Returns an error if there are no more values or
	// unmarshaler encounters an error
	ReadNext() (K, error)

	// More Returns whether there is next value
	More() bool
}
