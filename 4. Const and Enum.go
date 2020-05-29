package main

import "fmt"

// Used for enumerated expressions
const (
	_   = iota
	six = iota + 5
	seven
	eight
)

func main() {
	/*
		Constants:
			Typed constants
			Untyped constants
		Enumerated constants
		Enumerated expressions
	*/

	// Typed constants
	fmt.Println("Typed constants:-:\n")
	const myConst int = 12
	fmt.Printf("The myConst is %v, %T\n", myConst, myConst)

	// Untyped constants
	fmt.Println("Untyped constants:-:\n")
	const myUnConst = 12
	fmt.Printf("The myConst is %v, %T\n", myUnConst, myUnConst)

	/*
		Points:
		 In Go we cannot assign values as constants which takes value during runtime
		 Like : const myConst int = math.Sin(1.43)
		 Inner constants overshadows outer constants
	*/

	// Enumerated constants
	fmt.Println("Enumerated constants:-:\n")

	const (
		a1 = iota
		b1 = iota
		c1 = iota
	)

	const (
		a2 = iota
		b2
		c2
	)

	fmt.Printf("%v, %T\n", a1, a1)
	fmt.Printf("%v, %T\n", b1, b1)
	fmt.Printf("%v, %T\n", c1, c1)
	fmt.Printf("%v, %T\n", a2, a2)
	fmt.Printf("%v, %T\n", b2, b2)
	fmt.Printf("%v, %T\n", c2, c2)

	// Enumerated expressions
	fmt.Println("Enumerated expressions:-:\n")

	fmt.Printf("%v, %T\n", six, six)
	fmt.Printf("%v, %T\n", seven, seven)
	fmt.Printf("%v, %T\n", eight, eight)

	/*
		Points:
		 If we dont care about the zero and dont want to give memory to it we can use _ = iota as the first one
	*/
}
