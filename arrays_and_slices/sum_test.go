package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of 5 items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum all slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{0, 9})
		want := []int{6, 9}
		assertSlicesEqual(t, got, want)
	})

	t.Run("safety sum empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{0, 9, 1})
		want := []int{0, 10}
		assertSlicesEqual(t, got, want)
	})

	t.Run("safety sum many empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{}, []int{})
		want := []int{0, 0, 0}
		assertSlicesEqual(t, got, want)
	})
}

func assertSlicesEqual(t testing.TB, got, want []int) {
	t.Helper()
	if !compareSlices(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
