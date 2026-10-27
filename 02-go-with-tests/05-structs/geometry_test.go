package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "should calculate the perimeter of a rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasArea: 40.0},
		{name: "should calculate the perimeter of a circle", shape: Circle{Radius: 6.0}, hasArea: 37.69911184307752},
		{name: "should calculate the perimeter of a triangle", shape: Triangle{Base: 4.0, Height: 5.0}, hasArea: 14.77032961426901},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasArea {
				t.Errorf("given %#v, got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "should calculate the area of a rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "should calculate the area of a circle", shape: Circle{Radius: 7.0}, hasArea: 153.93804002589985},
		{name: "should calculate the area of a triangle", shape: Triangle{Base: 12.0, Height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("given %#v, got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
