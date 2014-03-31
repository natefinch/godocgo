package sub

// Constants are sorted to the top of the code documentation.
//
// Formatting
//
// Because it is declared on its own and not in a const() block, Applename gets
// full formatting support.
//
// Note that the declarations are sorted alphabetically, not based on what order
// they exist in the code.
const AppleName = "apple"

const (
	// PearName is an exported constant used to create pear foods.
	// Constants and variables enclosed in a block like this one get printed out
	// as-is.
	PearName = "pear"
)

// Apple is a global variable in the package.
//
// Formatting
//
// Because apple is declared on its own, rather in a block of var (), its
// documentation gets full formatting support.
var Apple = NewFood(AppleName)

var (
	// Pear is another global variable in the package.
	//
	// Because it's in a var block, it just gets plaintext docs.  Poor pear.
	Pear = NewFood(PearName)
)

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
