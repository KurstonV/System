package main

import (
	"fmt"
	"strconv"
)

// create an interfaces using the interface type

type Shape interface {
	Area() float64
}

// Creating the cirle type
type Circle struct {
	Radius float64
}

// Creating the rectangle type

type Rectangle struct {
	Length float64
	Width  float64
}

type Square struct {
	Length float64
	Width  float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Parallelogram struct {
	Base   float64
	Height float64
}

// Area method to the circle type
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Area method to the rectangle Type
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Area method to the Square
func (s Square) Area() float64 {
	return s.Length * s.Width
}

// Area method to the Triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

// Area method to the parallelogram
func (p Parallelogram) Area() float64 {
	return p.Base * p.Height
}

// Add a function to print a circle
func PrintCircleArea(c Circle) {
	fmt.Printf("Area: %.2f\n", c.Area())
}

// Add a function to Print a Rectangle
func PrintRectangleArea(r Rectangle) {
	fmt.Printf("Area: %.2f\n", r.Area())
}

// Add a function to Print a Square
func PrintSquareArea(s Square) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

// Add a function to Print a triangle
func PrintTriangleArea(t Triangle) {
	fmt.Printf("Area: %.2f\n", t.Area())
}

// Add a fucntion to Print a parallelogram
func PrintParallelogramArea(p Parallelogram) {
	fmt.Printf("Area: %.2f\n", p.Area())
}

func main() {
	var Radiusvalue string
	fmt.Println("Please Enter a value for the radius: ")
	fmt.Scanln(&Radiusvalue)

	UserRadius, err := strconv.ParseFloat(Radiusvalue, 64)

	if err != nil {
		fmt.Println("Your value is invalid!!")
	}

	circle := Circle{
		Radius: UserRadius,
	}
	rectangle := Rectangle{
		Length: 6,
		Width:  4,
	}
	square := Square{
		Length: 9,
		Width:  9,
	}
	triangle := Triangle{
		Base:   12,
		Height: 20,
	}
	parallelogram := Parallelogram{
		Base:   10,
		Height: 10,
	}

	PrintCircleArea(circle)
	PrintRectangleArea(rectangle)
	PrintSquareArea(square)
	PrintTriangleArea(triangle)
	PrintParallelogramArea(parallelogram)
}
