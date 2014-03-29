package sub

// You can embed a runnable example in your package overview by creating the
// function func Example() in a test file (either using the package
// <package>_test declaration or in a file that ends in _test.go).
//
// If the contents of a test file contain just a single example function, code
// outside the example function, and no other test or benchmark functions, the
// entire file will be used as the example.  In that case, the example function
// will display as func main() as if it were a self-contained executable
// package.
//
// If you have multiple example functions in the same file or examples and tests
// in the same file, only the actual example function will be displayed in the
// documentation (which can be handy if you want the example to be shorter)
import (
	"fmt"
)

// The function below is actually written as "func Example()" in the code, but
// because this is a full-file example, godoc translates it to func main() to
// make it into a full go command.

// This text is the comment on the Example() function itself, and is
// displayed ahead of the actual example code in the documentation.
func Example() {
	hello()
}

// Because we have a function external to the example, and no other test or
// example functions in the file, we get to use the whole file as the example.
func hello() {
	fmt.Println("Hello example!")
}
