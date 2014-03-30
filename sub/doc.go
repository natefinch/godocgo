// Package sub contains the rest of the article.
//
// This Text
//
// In the files of a library package, you can attach a comment to the package
// name, which will be shown in the overview for that package's documentation.
// Just like for command docuimentation, the convention is to use a file called
// doc.go for your package documentation, though it is not necessary.  If you
// have comments on the package name in multiple files, they'll be concatenated,
// but generally it's best to just have comments in one file.
//
// Just as with the command documentation, the first sentence in the
// documentation is used as the short description of the code in the package
// index. By convention, it should start with the words "Package <packagename>".
//
// Examples
//
// One of the coolest things about godoc is that you can have documentation that
// is also a test.  These are called examples, and they are specially named
// functions in test files (files that end in _test.go) that have their code
// automatically embedded in the documentation for a package, type, method, or
// function.
//
// To write an example function on the package as a whole, simply put a function
// in a test file called func Example().  For functions and types, the function
// name should be ExampleName(), where Name is the name of the type or function.
// To make an example for a method on a type, the function name should be
// ExampleType_Method, where type is the name of the type and method is the name
// of the method on the type.
//
// If you need more than one example for any of these, you can make additional
// functions with the same name, and a suffix of _sometext, where sometext is
// whatever lowercase text you want.
//
// What gets displayed
//
// If the contents of a test file contain just a single example function, code
// outside the example function, and no other test or benchmark functions, the
// entire file will be used as the example.  In that case, the example function
// will display as func main() as if it were a self-contained executable
// package.  Traditionally, full file examples are placed in a file called
// example_test.go, though this is not required.
//
// If you have multiple example functions in the same file or examples and other
// tests in the same file, only the actual example function and code internal to
// it will be displayed in the documentation (which can be handy if you want the
// example to be shorter).
//
// Examples as tests
//
// Examples aren't just for documentation, they're actual tests too. During a
// run of "go test", the code in an example is run.  This makes sure your
// examples will always compile.  In addition, if the example outputs to stdout,
// you can actually test the output, and have a failing test if the output
// doesn't pass. To do this, simply put a comment inside the function, at the
// end, with the first line as "// output:". Further lines will be stripped of
// their comment prefix, trimmed of whitespace, and matched against the
// example's output (line breaks in the comments match "\n" from the output of
// the example)
//
// Package example
//
// The example below is an example of a whole-file example, and also of an
// example on the package itself.  Expand the example to read more about whole
// file examples.  Note that the example is placed here, because the name of the
// example function is "Example()", and thus is an example about the entire
// package.
package sub
