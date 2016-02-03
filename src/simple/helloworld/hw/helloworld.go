package hw

import "fmt"

// HelloWorld greets you
// what specifies an additional message
func HelloWorld(what string) string {
	return fmt.Sprintf("Hello World %s", what)
}
