package sub

// Food is a type to show how type documentation works.
type Food struct {
	// Name holds the common name of the food.
	Name string

	// weight is an unexported field that won't show up in documentation.
	weight int
}

// NewFood creates a new Food with the given name and a calculated weight.
func NewFood(name string) Food {
	return Food{name, len(name)}
}

// Weigh reports the weight of the food.
func (f Food) Weigh() int {
	return f.weight
}
