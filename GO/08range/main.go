package main

import "fmt"

func main() {
	arr1 := []int{4, 3, 2, 1}
	for index, val := range arr1 { //use "_" instead of index or val to discard it
		fmt.Println(index, val)
	}
}
