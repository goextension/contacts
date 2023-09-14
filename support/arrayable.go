package support

type Arrayable interface {

	// ToArray Convert the struct to its Array representation.
	ToArray() []any
}
