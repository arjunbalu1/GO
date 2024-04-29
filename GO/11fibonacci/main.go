package main

import "fmt"

func fibonacci() func() int {
	n := -1
	return func() int {
		n++
		if n == 0 {
			return 0
		} else if n == 1 {
			return 1
		} else {
			return ((n - 1) + (n - 2)) //wrong but get the point
		}
	}

}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
