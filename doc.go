// Copyright 2014 Nate Finch

// Command godocgo is a repo demonstrating effective Go documentation.
//
// Introduction
//
// I started to write this as a blog post, intending to link to a repo with
// examples, and then realized that there was no reason I couldn't write the
// whole post as godoc on the repo, so here it is.  This repo is an example of
// good practices for documenting Go code, and also serves as a working example
// of "hey, how do I get something like that to show up in my godoc?"
//
// You can fork this repo in github, or just go get
// github.com/natefinch/godocgo, and have it as a reference locally on your hard
// drive.
//
// To view the godoc on your own machine, just run
//
//   go get code.google.com/p/go.tools/cmd/godoc
//   godoc -http=:8080 -play
//
// Then you can open a web browser and go to
// http://localhost:8080/pkg/github.com/natefinch/godocgo/
//
// Alternatively, you can view this on the excellent godoc.org, by going to
// http://godoc.org/github.com/natefinch/godocgo (note that godoc.org has
// slighty different styling than the godoc tool).
//
// Godoc is Awesome
//
// One of my favorite things about the Go programming language is godoc, the
// auto-generated documentation for your code.  If you're coming from javadoc or
// similar tools, you're probably thinking one of two things (possibly both):
//
//   1) meh, we have that
//   2) gah! no!
//
// An understandable reaction, but let me show you why godoc is different, and
// how to leverage all its features.
//
// First off, godoc comments are designed from the start to be read in the code.
// We all know that the code will be read a lot more often than some html
// documentation out on the web, so godoc makes sure to keep its documentation
// legible for the coder in their text editor. There is not a single angle
// bracket to be found.  Simple conventions replace the html/xml tags you often
// see in other documentation standards.
//
// Formatting
//
// The line above above this is heading. It is simply a single line without
// punctuation with paragraphs before and after it.  As you can see, it gets
// special formatting.  If you're viewing this using godoc, it will be included
// in the table of contents at the top of the file.  If you're viewing this on
// godoc.org, there is, sadly, no table of contents.
//
// Below this is pre-formatted text.  All that is required for it to be rendered
// in that way is that it be indented from the surrounding comments (you can use
// tabs or spaces, whatever you like).  This is handy for showing code snippets
// that you don't want godoc to auto-wrap, the way it does normal text.
//
//   a := []string{"foo", "bar"}
//   for _, b := range a {
//   	fmt.Println(b)
//   }
//
// Html links will be auto linked without you having to do anything, like so:
// http://golang.org
//
// That's it, that's all the formatting you get.  But it's enough, and I think
// we can all be glad to focus on the content and not the formatting of our
// documentation.
//
// Attaching Comments
//
// In Go, any comment which immediately precedes Go code is considered a comment
// about that code, and godoc will associate the two in its documentation.
//
//   // Foo is a function that deserves a comment.  This comment becomes the
//   // godoc comment for Foo.
//   func Foo() {}
//
// By corollary, any comment that is *not* on a line before a line of code, is
// not picked up as documentation. Note that this is important when writing
// things like build tags and copyright notices at the top of files.  Always
// make sure you have a space between those comments and the package
// declaration.
//
// This Text
//
// The text you are reading is long form documentation on the main package of a
// Go executable.  The Go authors were aware that you might need a really long
// comment on your application, so they devoted an entire file to it - doc.go.
// The contents of the file are just a really long comment immediately followed
// by the package declaration "package main".  You can also have code in doc.go,
// but traditionally it only contains documentation.
//
// The first sentence in this file (up to the first period) will be the label
// next to the link in godoc's index.  This should be a short description of
// what the command does, and traditionally starts with "Command <name-of-
// command>".
//
// Note that doc.go only applies to the main package of applications, not for
// library packages.
//
// To read more of the article, and explore more of what godoc has to offer,
// click on the link to the package named "sub" below.
package main
