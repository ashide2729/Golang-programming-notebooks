package main

import (
	"fmt"
	"math"
)

// define an interface
type Shape interface {
	area() float64
}

// define a circle
type Circle struct {
	x, y, radius float64
}

// define a rectangle
type Square struct {
	side float64
}

// define a method for circle (implementation of Shape.area())
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// define a method for square (implementation of Shape.area())
func (sq Square) area() float64 {
	return sq.side * sq.side
}

// define a method for shape
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {

	fmt.Println("###Interfaces:-:\n")

	circle := Circle{x: 0, y: 0, radius: 5}
	square := Square{side: 10}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Square area: %f\n", getArea(square))
}
