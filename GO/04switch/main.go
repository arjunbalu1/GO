package main

import "fmt"

func main() {
	fmt.Println("what is 420 + 69?")
	var ans int
	fmt.Scanln(&ans)
	switch ans {
	case 1337:
		fmt.Println("LEET")
	case 420:
		fmt.Println("dank")
	case 69:
		fmt.Println("nice")
	case 489:
		fmt.Println("Correct")
	default:
		fmt.Println("Wrong")
	}
}
