package main

import "fmt"

func main() {

	mathMarks := [3]int{97, 98, 99}
	englishMarks := []int{97, 98, 99}
	scienceMarks := [...]int{97, 98, 99}
	var subjects [3]string
	fmt.Printf("Math Marks: %v, %T\n", mathMarks, mathMarks)
	fmt.Printf("English Marks: %v, %T\n", englishMarks, englishMarks)
	fmt.Printf("Science Marks: %v, %T\n", scienceMarks, scienceMarks)
	fmt.Printf("Subjects: %v, %T\n", subjects, subjects)
	subjects[0] = "Maths"
	subjects[1] = "English"
	subjects[2] = "Science"
	fmt.Printf("Subjects: %v, %T\n", subjects, subjects)

	// 2-D array
	var marksArray [3][3]int
	marksArray[0] = mathMarks
	marksArray[1] = [3]int{95, 96, 97}
	marksArray[2] = scienceMarks
	fmt.Printf("2-D Marks array: %v, %T\n\n", marksArray, marksArray)

	// Array copy vs pointers
	myArray := [...]int{1, 2, 3, 4, 5}
	myArrayCopy := myArray
	myArrayCopy[0] = 5
	fmt.Printf("The myArray is: %v\n", myArray)
	fmt.Printf("The myArrayCopy is: %v\n", myArrayCopy)

	myArrayPointerCopy := &myArray
	myArrayPointerCopy[0] = 5
	fmt.Printf("The myArray is: %v\n", myArray)
	fmt.Printf("The myArrayPointerCopy is: %v\n", myArrayPointerCopy)

	myArraySlice := []int{1, 2, 3, 4, 5}
	myArraySliceCopy := myArraySlice
	myArraySliceCopy[0] = 5
	fmt.Printf("The myArraySlice is: %v\n", myArraySlice)
	fmt.Printf("The myArraySliceCopy is: %v\n", myArraySliceCopy)

}
