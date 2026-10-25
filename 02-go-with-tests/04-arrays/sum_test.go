package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should sum a collection of N integers", func(t *testing.T) {
		numbers := []int{4, 5, 6, 7}

		got := Sum(numbers)
		want := 22

		if got != want {
			t.Errorf("got %#v want %#v given %#v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should sum a collection of N slices, returning N results", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5, 6, 7}, []int{8, 9}, []int{0, 0, 0})
		want := []int{6, 22, 17, 0}

		if !slices.Equal(got, want) {
			t.Errorf("got %#v want %#v", got, want)
		}
	})

	t.Run("should return N zeroes for a collection of N empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{}, []int{})
		want := []int{0, 0, 0}

		if !slices.Equal(got, want) {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("should sum a collection of N slices excluding the first item of each, and returning N results", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4, 5}, []int{0, 0})
		want := []int{2, 9, 0}

		if !slices.Equal(got, want) {
			t.Errorf("got %#v want %#v", got, want)
		}
	})

	t.Run("should return N zeroes for a collection of N empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{}, []int{})
		want := []int{0, 0}

		if !slices.Equal(got, want) {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}
