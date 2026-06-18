package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2.0) * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.Radius * math.Pi * 2
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func PrintPerimeter(s Shape) {
	fmt.Println(s.Perimeter())
}

func main() {
	r := Rectangle{
		Width:  15.0,
		Height: 25.0,
	}

	c := Circle{
		Radius: 4.0,
	}

	PrintArea(r)
	PrintPerimeter(c)
}
