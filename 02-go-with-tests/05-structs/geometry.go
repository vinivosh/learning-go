package main

import "math"

type Shape interface {
	// Perimeter returns the perimeter of the Shape.
	//
	// The case for triangles is simplified: it assumes an isosceles or equilateral triangle. A scalene triangle DOES NOT work!
	Perimeter() float64

	// Area returns the area of the Shape
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2.0
}

func (c Circle) Perimeter() float64 {
	return 2.0 * math.Pi * (c.Radius)
}

func (t Triangle) Perimeter() float64 {
	equalSideLen := math.Hypot(t.Base*0.5, t.Height)
	return t.Base + 2*equalSideLen
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
