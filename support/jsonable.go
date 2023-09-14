package support

type Jsonable interface {

	// ToJson Convert the struct to its JSON representation.
	ToJson() string
}
