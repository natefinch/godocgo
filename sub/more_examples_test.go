package sub

import (
	"fmt"
)

// This is an example function for a type.
func ExampleFood() {
	f := NewFood("apple")
	fmt.Println(f.Weight())
	// output:
	// 5
}

// Example on a function.
func ExampleNewFood() {
	f := NewFood("pear")
	fmt.Println(f.Name)
	// output:
	// pear
}

// Example on a method.  Note that if you run go test, this test will fail,
// because the output doesn't match.
func ExampleFood_Weight() {
	f := NewFood("pineapple")
	fmt.Println(f.Weight())
	// output:
	// 10
}
