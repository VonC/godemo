package main

import "fmt"

func main() {
	HelloWorld("in go")
}

// HelloWorld greets you
// what specifies an additional message
func HelloWorld(what string) {
	fmt.Printf("Hello World %s\n", what)
}
