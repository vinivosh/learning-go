package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("should sum a collection of 5 integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %#v want %#v given %#v", got, want, numbers)
		}
	})

	t.Run("should sum a collection of N integers", func(t *testing.T) {
		numbers := []int{4, 5, 6, 7}

		got := Sum(numbers)
		want := 22

		if got != want {
			t.Errorf("got %#v want %#v given %#v", got, want, numbers)
		}
	})
}
