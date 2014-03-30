package sub_test

import (
	"fmt"

	"github.com/natefinch/godocgo/sub"
)

// This is an example function for a type.  To make an example show up on a
// type, just put a function in a testfile called Example<Typename>(), for
// example, this function is func ExampleFood() in more_examples_test.go.  If
// you need more than one example function for a type, function, or method, just
// make more functions, except give them a unique lowercase suffix after an
// underscore, like ExampleFood_more().
func ExampleFood() {
	f := sub.NewFood("apple")
	fmt.Println(f.Weigh())
	// output:
	// 5
}
