package main

import (
	"fmt"
)

// BUG(natefinch) This is a bug description.  Bugs are an exception to the rule
// about only attached comments getting displayed in godoc.  Simply make a
// comment that starts with BUG(foo) where foo is the name of the person
// responsible for the bug, and it'll show up in the list of bugs in godoc.  You
// can click on the link at the beginning of the bug description to go right to
// the place where the BUG comment is located in the code.

func main() {
	fmt.Println("Hello world!")
}
