package main

import "testing"

func TestShape(t *testing.T) {
	t.Run("test area", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{10.0, 10.0}, 100.0},
			{Circle{10.0}, 314.1592653589793},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()
			assertFloatEquals(t, got, tt.want)
		}
	})

	t.Run("test perimeter", func(t *testing.T) {
		perimeterTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{10.0, 10.0}, 40.0},
			{Circle{10.0}, 62.83185307179586},
		}

		for _, tt := range perimeterTests {
			got := tt.shape.Perimeter()
			assertFloatEquals(t, got, tt.want)
		}
	})
}

func assertFloatEquals(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
