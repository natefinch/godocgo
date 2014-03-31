// Package sub contains the rest of the article.
//
// This Text
//
// In the files of a library package, you can attach a comment to the package
// name, which will be shown in the overview for that package's documentation.
// Just like for command documentation, the convention is to use a file called
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
// One of the coolest things about godoc is that you can have code snippets
// embedded in your documentation that get run with go test. These are called
// examples, and they are specially named functions in test files (files that
// end in _test.go) that have their code automatically embedded in the
// documentation for a package, type, method, or function.
//
// To write an example function on the package as a whole, simply put a function
// in a test file called func Example().  For functions and types, the function
// name should be ExampleName(), where Name is the name of the type or function.
// To make an example for a method on a type, the function name should be
// ExampleType_Method, where Type is the name of the type and Method is the name
// of the method on the type.
//
// If you need more than one example for any of these, you can make additional
// functions with the same name, and a suffix of _sometext, where sometext is
// whatever lowercase text you want.  The suffix will be shown in parentheses
// after the Example label.
//
// What gets displayed
//
// If the contents of a test file contain a single example function, code
// outside the example function, and no other test, example, or benchmark
// functions, the entire file will be used as the example.  Traditionally, full
// file examples are placed in a file called example_test.go, though this is not
// required.
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
// doesn't pass. To do this, simply put a comment at the end of the function,
// with the first line as "output:". Further lines will be stripped of their
// comment prefix, trimmed of whitespace, and matched against the example's
// output (line breaks in the comments match "\n" from the output of the
// example code).
//
// In the code, examples look like this:
//
//	func ExampleFood() {
//		f := sub.NewFood("apple")
//		fmt.Println(f.Weight())
//		// output:
//		// 5
//	}
//
// Runnable Examples
//
// Some of your examples may be able to be run right from the documentation,
// using the same technology that play.golang.org uses (in fact, godoc.org's
// support for this actually sends you to play.golang.org).  To see this on your
// local godoc, start godoc with the -play option.  However, be aware that the
// use of this option is limited.
//
// In order for an example to be runnable, it can only reference code in the
// same file as it, or code in the standard library.  This ends up being quite
// limiting, because you can't actually use any of the types or functions
// written in your package.
//
// Package example
//
// The example below is a whole-file example, and shows how to embed an example
// in the package documentation. Expand the example to read more about whole
// file examples.  Note that the example is placed here, because the name of the
// example function is "Example()", and thus is an example about the entire
// package.
package sub
