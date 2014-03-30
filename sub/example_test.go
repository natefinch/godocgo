// This statement is actually "package sub" in the file, but it gets translated
// to "package main" for the example, to make it a fully self-contained runnable
// file.
package sub

import (
	"fmt"
)

// The function below is actually written as "func Example()" in the code, but
// because this is a full-file example, godoc translates it to func main() to
// make it into a full go command.

// This text is the comment on the Example() function itself, and is displayed
// ahead of the actual example code in the documentation.  Note that unlike docs
// elsewhere, docs on examples are plaintext only, so be brief.
func Example() {
	hello()
	// output:
	// Hello example!
}

// Because we have code external to the example, and no other test or
// example functions in the file, godoc uses the whole file as the example.
func hello() {
	fmt.Println("Hello example!")
}
