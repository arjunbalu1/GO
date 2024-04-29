package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "Alice", age: 25}
	ptr := &p

	// Print the pointer address
	fmt.Printf("Pointer address: %p", ptr)
}
