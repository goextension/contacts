package support

type Enumerable[T any] interface {

	// Pluck Get the values of a given key.
	Pluck(value string, column string) T
}
