package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func add1() int {
	i := 0
	i++
	return i
}

func main() {

	nextInt := intSeq()
	fmt.Println("calling function - nextInt x3")
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println("calling function - newInt x3")
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())

	fmt.Println("calling functin without a closure x3")
	fmt.Println(add1())
	fmt.Println(add1())
	fmt.Println(add1())
}
