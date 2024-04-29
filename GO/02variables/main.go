package main

import "fmt"

const ConstantVar1 string = "if first letter of variable is capital, then it is public (accessible by other packages)"

//or can be var

func main() {
	var name string = "Arjun"
	fmt.Println(name)
	fmt.Printf("variable is of type: %T \n", name)

	var isLegend bool = true
	fmt.Println(isLegend)
	fmt.Printf("variable is of type: %T \n", isLegend)

	var smallVal uint8 = 255 //unsigned int
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.1283912308293823
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var intVar int
	fmt.Println(intVar)
	fmt.Printf("Variable is of type: %T \n", intVar)

	var stringVar string
	fmt.Println(stringVar)
	fmt.Printf("Variable is of type: %T\n", stringVar)

	//implicit type
	var stringVar2 = "i didn't declare this as a string"
	fmt.Println(stringVar2)
	fmt.Printf("Variable is of type: %T \n", stringVar2)

	//no var style
	testVar := 69.420 //can only use this style inside a method (can't use as global)
	fmt.Println(testVar)
	fmt.Printf("Variable is of type: %T \n", testVar)

	fmt.Println(ConstantVar1)
	fmt.Printf("Variable is of type: %T \n", ConstantVar1)
}
