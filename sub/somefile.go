package sub

import "fmt"

// BUG(natefinch) This is a bug description. Bugs are an exception to the rule
// about only attached comments getting displayed in godoc.  Simply make a
// comment that starts with BUG(foo) where foo is the name of the person
// responsible for the bug, and it'll show up in the list of bugs in godoc.
// Note that bug descriptions are plaintext only.  You can click on the link at
// the beginning of the bug description to go right to the place where the BUG
// comment is located in the code.

// BUG(godoc.org) Note that bugs in the command (application) package do not
// show up on godoc.org.  This appears to be a bug in godoc.org, as they do show
// up in local godoc.

// Constants are sorted to the top of the code documentation.
//
// Formatting
//
// Because it is declared on its own and not in a const() block, AppleName gets
// full formatting support.
const AppleName = "apple"

// This text is a doc comment on the const block as a whole.  Unlike the
// comments inside the block, this text has full formatting.
//
// This can be a nice way to write a single comment about a group of variables.
const (
	// PearName is an exported constant used to create pear foods. Docs on
	// constants and variables enclosed in a block like this one get printed out
	// as-is.
	PearName = "pear"

	// PineappleName is the name of a pineapple.
	PineappleName = "pineapple"
)

// Apple is a global variable in the package.
//
// Because apple is declared on its own, rather in a block of var (), its
// documentation gets full formatting support.
var Apple = NewFood(AppleName)

// This is a block of variables with a formatted comment.
//
// Pretty nice.
var (
	// Pear is another global variable in the package.
	Pear = NewFood(PearName)
)

// Favourite is a global variable in the package.
//
// Because it is declared with a type defined within the package, it is
// sorted with the type definition in the documentation.
var Favourite Food = Pear

// Food is a type to show how type documentation works.
//
// Formatting
//
// You always get full formatting on type, function, and method docs.
//
// Below is the auto-embedded example from the function ExampleFood in
// more_examples_test.go.
type Food struct {
	// Name holds the common name of the food.
	//
	// Alas, no formatting here.
	Name string

	// weight is an unexported field that won't show up in documentation.
	weight int
}

// NewFood creates a new Food with the given name and a calculated weight.
//
// Godoc sorts "constructor" style functions like this, that return a type
// defined in the package, under the type itself.
//
// Note that you do not need to (and should not) make separate lines describing
// each parameter to the function.  Simply refer to the parameters by their name
// when you need to reference them in the rest of the documentation.
func NewFood(name string) Food {
	return Food{name, len(name)}
}

// Weight reports the weight of the food in letters.
//
// When decribing a function without parameters that simply returns data, it is
// traditional to write the doc saying the function reports foo, rather than
// *returns* foo.
//
// The example below intentionally fails during go test, to show you what that
// looks like.  The test will report:
//
//   --- FAIL: ExampleFood_Weight (4.933us)
//   got:
//   9
//   want:
//   10
//   FAIL
//   exit status 1
func (f Food) Weight() int {
	return f.weight
}

// Zap consumes the food and prints out a nasty message.
//
// This is just another top level function to show how the alphabetical sorting
// works.
func Zap(f Food) {
	fmt.Printf("I hate %ss!", f.Name)

}

// Eat consumes the food and prints out a nice message.
//
// This is an example of a function that isn't a "constructor" style function,
// so it will be sorted at the top level of the documentation.
//
// Functions get sorted before types.
func Eat(f Food) {
	fmt.Printf("I love %ss!", f.Name)
}

// Weight is a type to show how type-associated const and var documentation works.
type Weight int

// Constant declarations are sorted with their type definition when all the
// constant types match.
//
// This works with elided type definition using iota, but does not work when
// the constant type is deduced from the constant expression.
const (
	Small Weight = 1  // This much fruit is not enough.
	Big   Weight = 20 // This is way too much fruit.
)

// Variable declarations of a specific type are also associated with their type
// declaration in the documentation.
var WeightPreference Weight
