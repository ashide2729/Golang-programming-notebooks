package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("###Maps:-:\n")
	classMarksMap := map[string]int{
		"Alan":   98,
		"Josh":   89,
		"Danny":  98,
		"Dexter": 77,
	}
	fmt.Printf("The classMarksMap are: %v, %T\n", classMarksMap, classMarksMap)

	classMarksMapAnother := make(map[string]int)
	classMarksMapAnother = map[string]int{
		"Alan":   98,
		"Josh":   89,
		"Danny":  98,
		"Dexter": 77,
	}
	fmt.Printf("The classMarksMapAnother are: %v, %T\n", classMarksMapAnother, classMarksMapAnother)

	println("The marks of Alan are:", classMarksMap["Alan"])
	classMarksMap["Alan"] = 0
	println("The marks of Alan after change are:", classMarksMap["Alan"])

	// Maps in go provide a value of zero even if the key doesn't exist so it is difficult to find
	// if element value is zero or the element is not present

	println("The marks of RandomGuys after change are:", classMarksMap["RandomGuys"])

	println("The marks of Alen after change are:", classMarksMap["Alen"])

	val, ok := classMarksMap["RandomGuys"]
	println("The marks of RandomGuys after change are:", val, ok)
	_, ok1 := classMarksMap["Alan"]
	println("The marks of Alan after change are:", ok1)

	fmt.Println("\n###Structs:-:\n")
	type Student struct {
		id           int
		name         string
		fresher      bool
		certificates []string
	}

	Student1 := Student{
		id:      1,
		name:    "A",
		fresher: true,
		certificates: []string{
			"A", "B", "C",
		},
	}

	fmt.Println("The student1 is:", Student1)
	fmt.Println("The name of student1 is:", Student1.name)
	fmt.Println("The top two certificates of student1 are:", Student1.certificates[:2])

	type Bird struct {
		name string
	}
	// Go doesn't have inheritance but it is called composition via embedding
	type Parrot struct {
		Bird
		featherColour string
	}
	anotherParrot := Parrot{}
	anotherParrot.name = "parrot1"
	anotherParrot.featherColour = "green"
	fmt.Println("The anotherParrot is:", anotherParrot)

	type TagStruct struct {
		name string `required max:"10"`
		age  int
	}
	tag := reflect.TypeOf(TagStruct{})
	conditions, _ := tag.FieldByName("name")
	fmt.Println("The tag in the name field is:", conditions.Tag)
}
