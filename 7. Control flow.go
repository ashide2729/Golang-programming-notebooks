package main

func main() {
	println("###If-else:-:\n")

	if true {
		println("I am true!")
	}

	if a := 6 > 5; a {
		println("6 is greater than 5")
	}

	value := 12
	if value > 11 && value < 13 {
		println("The value is 12")
	} else {
		println("The value is not 12")
	}

	println("\n###Switch:-:\n")

	switch 4 > 3 {
	case true:
		println("Great math")
	case false:
		println("Wrong expression")
	default:
	}

	switch switchReturn() {
	case "A", "C":
		println("A or C")
	case "B", "D":
		println("B or D")
	default:
		println("Not A, B, C, D")
	}

	val := 1
	switch {
	case val > 1:
		println("Val greater than 1")
	case val < 1:
		println("Val less than 1")
	default:
		println("Val is 1")
	}
}
func switchReturn() string {
	return "B"
}
