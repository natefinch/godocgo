// Package sub contains the rest of the article.
//
// This Text
//
// In the files of a library package, you can attach a comment to the package
// name, which will be shown in the overview for that package's documentation.
// Note that unlike doc.go for command documentation, there is no special
// filename you need to use.  Just use whatever filename you like.  If you have
// comments on the package name in multiple files, they'll be concatenated, but
// generally it's best to just have comments in one file.
//
// Just as with the command documentation, the first sentence in the
// documentation is used as the short description of the code in the package
// index. By convention, it should start with the words "Package <packagename>".
//
// Expand the Example below to learn how to embed example code in your package
// overview.
package sub

// Food is a type to show how type documentation works.
type Food struct {
	// Name holds the common name of the food.
	Name string

	// weight is an unexported field that won't show up in documentation.
	weight int
}

// NewFood creates a new Food with the given name and a calculated weight.
func NewFood(name string) Food {
	return Food{name, len(name)}
}

// Weigh reports the weight of the food.
func (f Food) Weigh() int {
	return f.weight
}
