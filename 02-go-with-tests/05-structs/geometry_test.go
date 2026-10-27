package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"should calculate the perimeter of a rectangle", Rectangle{10.0, 10.0}, 40.0},
		{"should calculate the perimeter of a circle", Circle{6.0}, 37.69911184307752},
		{"should calculate the perimeter of a triangle", Triangle{4.0, 5.0}, 14.77032961426901},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"should calculate the area of a rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"should calculate the area of a circle", Circle{7.0}, 153.93804002589985},
		{"should calculate the area of a triangle", Triangle{12.0, 6.0}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}
