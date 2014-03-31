package sub

import (
	"fmt"
)

// This is an example function for a type.  This line comes from the comment on
// the example function.  See how the code is is displayed and then the output?
// Note that the output is just printing what you put under the output: comment
// in the function. It's not actually running the code to get that output.
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
// match the comment (the actual output is 9).  Note that the docs still say the
// output is 10, because that's what the comment says in the code.
func ExampleFood_Weight() {
	f := NewFood("pineapple")
	fmt.Println(f.Weight())
	// output:
	// 10
}
