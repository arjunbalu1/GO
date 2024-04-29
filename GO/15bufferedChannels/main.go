package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	ch <- 4
	//ch <- 5 //deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
