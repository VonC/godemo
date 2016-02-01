// Godoc example on package main
//
// More examples at:
//
// - https://github.com/natefinch/godocgo
// - https://godoc.org/github.com/fluhus/godoc-tricks
package helloworld

import "fmt"

// HelloWorld greets you.
// what specifies an additional message
func HelloWorld(what string) {
	fmt.Printf("Hello World %s\n", what)
}
