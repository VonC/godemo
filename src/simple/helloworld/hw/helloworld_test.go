package hw_test

import (
	"fmt"
	"simple/helloworld/hw"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	thw := hw.HelloWorld("as a test")
	if thw != "Hello World as a test" {
		t.Fatalf("Wrong HelloWorld: '%s'", thw)
	}
}

// ExampleHelloWord shows how to call HelloWorld()
func ExampleHelloWorld() {
	fmt.Println(hw.HelloWorld("as an example"))
	// Output: Hello World as an example
}
