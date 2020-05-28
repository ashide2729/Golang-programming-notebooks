package main

import (
	"fmt"
	"strconv"
)

// Variables can also be declared outside of main but ":=" cannot be used here
var myBoolVariable bool = true

// Variables with LOWERCASE name at package level are accessible to this file
// Like :   var i int = 42
// Variables with UPPERCASE name at package level are accessible to other files also which import this
// Like :   var I int = 42

func main() {
	fmt.Println("Declaring variables\n")
	// Declaring a variable is writing the way we read it "variable name type"
	var myIntVariable int
	myIntVariable = 37
	var myStrVariable string = "Hello, I am a string."
	// Using const keyword makes the variables constact throughout and cannot be changed later
	const x = 3
	// Variables can be declared with defining type where go automatically finds the datatype
	myAnotherVariable := 42.02
	fmt.Println("The value of myIntVariable is :", myIntVariable)
	fmt.Println("The value of myStrVariable is :", myStrVariable)
	fmt.Println("The value of myBoolVariable is :", myBoolVariable)
	fmt.Println("The value of myAnotherVariable is :", myAnotherVariable)
	fmt.Printf("The datatype of myAnothervariable is: %T\n\n\n", myAnotherVariable)

	/*
		Points for variables in Go:
			Cannot redeclare variables
			All variables must be used else error will be thrown
			Keep in mind the case of variable name based on use case
			Longer names for longer lives and vice versa
	*/

	// Type conversion
	fmt.Println("Type conversion\n")
	var a = 50
	fmt.Printf("%v, %T\n", a, a)
	b := float32(a)
	fmt.Printf("%v, %T\n", b, b)
	var c string
	c = string(a)
	// This will print value of 50 in unicode because in go string is an alias for byte
	fmt.Printf("%v, %T\n", c, c)
	c = strconv.Itoa(a)
	// This will convert int to string
	fmt.Printf("%v, %T\n", c, c)

	/*
		Points for type conversion in Go:
			No implicit type conversion i.e. float32 to int because of information loss
			Use strConv for string conversions
	*/

}
