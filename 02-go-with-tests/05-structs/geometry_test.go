package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("should calculate the area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.4f want %.4f", got, want)
		}
	})

	t.Run("should calculate the area of a circle", func(t *testing.T) {
		circle := Circle{6.0}
		got := circle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("should calculate the area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.4f want %.4f", got, want)
		}
	})

	t.Run("should calculate the area of a circle", func(t *testing.T) {
		circle := Circle{7.0}
		got := circle.Area()
		want := 21.9911485751285

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
