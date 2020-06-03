package main

import "fmt"

func main() {

	fmt.Println("###Pointers:-:\n")

	var a int = 50
	var pointerToA *int = &a
	println("The address to 50 for a:", &a, "The address to 50 for b:", pointerToA)
	println("The value at address for a:", a, "The value at address for b:", *pointerToA)

	fmt.Println("\n###Functions:-:\n")

	funcWithoutReturn()
	value := funcWithReturn()
	println("Value from funcWithReturn:", value)

	i, s := funcWithMultipleReturn()
	println("Values from funcWithMultipleReturn are:", s, i)

	funcWithParameters("Adam", 13)
	funcWithNParameters(1, 3, 4, 5)

	func() {
		println("I am from anonymous function")
	}()

	funcAsVariable := func() {
		println("I am from funcAsVariable")
	}
	funcAsVariable()

	fmt.Println("\n###Methods:-:\n")

	student := myStruct{
		"Lewis",
		18,
	}
	student.structMethod()

}

func funcWithoutReturn() {
	println("From function without return")
}

func funcWithReturn() int {
	println("From function with return")
	return 4
}

func funcWithMultipleReturn() (int, string) {
	println("From function with multiple return")
	return 4, "Charles"
}

func funcWithParameters(name string, age int) {
	println("Parameters from funcWithParameters are:", name, age)
}

func funcWithNParameters(numbers ...int) {
	println("The parameter from funcWithNParameters is:", numbers)
	fmt.Println("The parameter value from funcWithNParameters is: ", numbers)
}

type myStruct struct {
	name string
	age  int
}

func (str myStruct) structMethod() {
	fmt.Println("The values from method using struct as context are:", str.name, str.age)
}
