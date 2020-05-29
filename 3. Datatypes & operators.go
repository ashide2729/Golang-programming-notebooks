package main

import "fmt"

func main() {
	/*
		There are three Primitive datatypes:
			Boolean
			Numeric
				Integer
					Signed - 8, 16, 32, 64
					Unsigned - 8, 16, 32, 64
				Float
				Complex
			Text
	*/

	// Boolean datatypes
	fmt.Println("Boolean:-:\n")
	var myBool1 bool = true
	fmt.Println("The value of myBool1 is:", myBool1)
	myBool2 := 1 > 2
	fmt.Println("The value of myBool2 is:", myBool2)

	// Integer datatypes
	fmt.Println("\nInteger:-:\n")
	var mySignedInt int8 = -36
	fmt.Printf("%v, %T\n", mySignedInt, mySignedInt)
	var myUnsignedInt uint8 = 36
	fmt.Printf("%v, %T\n", myUnsignedInt, myUnsignedInt)

	// Float datatypes
	fmt.Println("\nFloat:-:\n")
	pi := 3.14
	pi = 1.2e12
	pi = 1.4E13
	fmt.Printf("The pi is : %v, %T\n", pi, pi)

	// Complex datatypes
	fmt.Println("\nComplex:-:\n")
	var eta complex64 = 1 - 1i
	var anotherEta complex128 = complex(10, 10)
	fmt.Printf("The eta is : %v, %T\n", eta, eta)
	fmt.Printf("The anortherEta is : %v, %T\n", anotherEta, anotherEta)
	fmt.Printf("The eta real part is : %v, %T\n", real(eta), real(eta))
	fmt.Printf("The eta imaginary part is : %v, %T\n", imag(eta), imag(eta))

	// String datatypes
	fmt.Println("\nString:-:\n")
	myString := "Hi ! I am a string."
	fmt.Printf("The myString is : %v, %T\n", myString, myString)
	// Strings can be trated as an array also with each element treated as a byte
	fmt.Printf("The myString 4th element is : %v, %T\n", myString[3], myString[3])
	fmt.Printf("The myString 4th element is : %v, %T\n", string(myString[3]), myString[3])

	// Rune datatypes
	// Runes are actually type alias for int32
	var r1 rune = 'a'
	r2 := '%'
	fmt.Printf("The rune1 is : %v, %T\n", r1, r1)
	fmt.Printf("The rune2 is : %v, %T\n", r2, r2)

	// Arithematic operations
	a := 10
	b := 3

	fmt.Println("\nArithematic Operations on a=10 & b=3:-:\n")

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Modulo:", a%b)

	// Bit operations

	fmt.Println("\nBit Operations on a=10 & b=3:-:\n")

	fmt.Println("AND:", a&b)
	fmt.Println("OR:", a|b)
	fmt.Println("XOR:", a^b)
	fmt.Println("NAND:", a&^b)
	fmt.Println("Left shift:", a<<3)
	fmt.Println("Right shift:", a>>3)

}
