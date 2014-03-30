package sub

// This is the actual contents of example_test.go.

import (
	"fmt"
)

// This text is the comment on the Example() function itself, and is displayed
// ahead of the actual example code in the documentation.  Note that unlike docs
// elsewhere, docs on examples are plaintext only, so be brief.
func Example() {
	hello()
	// output:
	// Hello example!
}

// Because we have code external to the example, and no other test,  or
// example functions in the file, godoc uses the whole file as the example.
func hello() {
	fmt.Println("Hello example!")
}
