package main

import "fmt"

func main() {
	fmt.Println("###For loop:-:\n")

	for i := 0; i < 3; i++ {
		fmt.Println("The value of i is: ", i)
	}

	for i, j := 1, 2; i < 9; i, j = i+2, j+2 {
		fmt.Println("The odd(i) and even(j) are: ", i, j)
	}

	fmt.Println("###While loop:-:\n")
	j := 0
	for j < 3 {
		println("From while loop", j)
		j++
	}

	fmt.Println("\n###Break and continue:-:\n")
	i := 0

	for {
		i++
		if i > 5 {
			println("Breaking at ", i)
			break
		}
		if i%2 == 0 {
			continue
		}
		println("The i value is:", i)

	}

	fmt.Println("\n###Labels:-:\n")

OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			println("From nested loop i+j:", i+j)
			if i+j == 3 {
				break OuterLoop
			}
		}
	}

	fmt.Println("\n###Loops with collections:-:\n")
	arr := []int{1, 2, 3}
	for k, v := range arr {
		fmt.Println("Index", k, "Value", v)
	}

}
