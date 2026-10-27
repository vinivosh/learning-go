package main

import "testing"

func CheckPerimeter(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Perimeter()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("should calculate the perimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 40.0

		CheckPerimeter(t, rectangle, want)
	})

	t.Run("should calculate the perimeter of a circle (a.k.a. its circumference)", func(t *testing.T) {
		circle := Circle{6.0}
		want := 37.69911184307752

		CheckPerimeter(t, circle, want)
	})
}

func CheckArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("should calculate the area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0

		CheckArea(t, rectangle, want)
	})

	t.Run("should calculate the area of a circle", func(t *testing.T) {
		circle := Circle{7.0}
		want := 153.93804002589985

		CheckArea(t, circle, want)
	})
}
