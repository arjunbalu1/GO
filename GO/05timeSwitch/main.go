package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	fmt.Println("Today is", today)
	fmt.Print("Saturday is: ")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
