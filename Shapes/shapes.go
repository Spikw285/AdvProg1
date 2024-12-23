package Shapes

import (
	"math"
)

// Shape interface with methods for calculating area and perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct with length and width
type Rectangle struct {
	Length float64
	Width  float64
}

// Circle struct with radius
type Circle struct {
	Radius float64
}

// Square struct with length (side)
type Square struct {
	Length float64
}

// Triangle struct with sides
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

// Implement Area() and Perimeter() methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Implement Area() and Perimeter() methods for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// Implement Area() and Perimeter() methods for Square
func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (s Square) Perimeter() float64 {
	return 4 * s.Length
}

// Implement Area() and Perimeter() methods for Triangle
func (t Triangle) Area() float64 {
	// Using Heron's formula for area
	p := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(p * (p - t.SideA) * (p - t.SideB) * (p - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}
