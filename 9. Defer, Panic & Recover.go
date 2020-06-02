package main

import "fmt"

func main() {
	fmt.Println("###Defer:-:\n")

	deferExample()

	fmt.Println("\n###Panic:-:\n")

	a, b := 1, 0
	panic("Divide by zero")
	answer := a / b

	println(answer)

	fmt.Println("\n###Recover:-:\n")
}

func deferExample() {
	println("A")
	defer println("B")
	println("C")
}
