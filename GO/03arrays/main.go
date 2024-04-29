package main

import "fmt"

func main() {
	arr1 := []int{5: 10, 8: 20} //initializing only specific elements
	fmt.Println("length of arr1 is:", len(arr1))
	fmt.Println("arr1 elements are:", arr1)
	fmt.Println("2nd element is:", arr1[1])
	fmt.Println("6th element is:", arr1[5])
	fmt.Println("changing value of 5th element...")
	arr1[4] = 69
	fmt.Println("new arr1 is:", arr1)
	fmt.Println("creating new slice (3:7)")
	slice1 := arr1[3:7]
	fmt.Println("slice1 is:", slice1)
	fmt.Println("length is:", len(slice1))
	fmt.Println("capacity is:", cap(slice1))
}
