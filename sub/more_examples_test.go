package sub

import (
	"fmt"
)

// This is an example function for a type.  The name of the example function is
// ExampleFood, which is why it's showing up here.
func ExampleFood() {
	f := NewFood("apple")
	fmt.Println(f.Weight())
	// output:
	// 5
}

// Example documentation on a function.
func ExampleNewFood() {
	f := NewFood("pear")
	fmt.Println(f.Name)
	// output:
	// pear
}

// Second example on the same function, the name of this example function is
// ExampleNewFood_aux.
func ExampleNewFood_aux() {
	f := NewFood("prickly pear")
	fmt.Println(f.Name)
	// output:
	// prickly pear
}

// Example on a method, the name of this example function is ExampleFood_Weight.
// Note that if you run go test, this test will fail, because the output doesn't
// match the comment (the actual output is 9).
func ExampleFood_Weight() {
	f := NewFood("pineapple")
	fmt.Println(f.Weight())
	// output:
	// 10
}
