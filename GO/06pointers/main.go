package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(p)
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	fmt.Println(*p)

	p = &j // point to j
	fmt.Println(*p)
	*p = *p / 37 // divide j through the pointer
	fmt.Println(*p)
	fmt.Println(j) // see the new value of j

	v := Vertex{1, 2}
	l := &v
	fmt.Printf("%p \n", l)
	l.X = 3
	fmt.Println(v)

}
