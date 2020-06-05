package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("###Routines:-:\n")
	msg := "Hi"
	func() {
		println("The message is:", msg)
	}()
	msg = "Bye"

	msg1 := "Hi"
	go func() {
		println("The message using goroutine is:", msg1)
	}()
	msg1 = "Bye"

	time.Sleep(10 * time.Millisecond)

	fmt.Println("\n###Channels:-:\n")

}
