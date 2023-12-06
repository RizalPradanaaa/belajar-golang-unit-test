package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		panic("Result not Hello Eko")
	}
}
