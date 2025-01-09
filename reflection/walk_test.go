package main

import "testing"

func TestWalk(t *testing.T) {
	t.Run("test walk calls fn", func(t *testing.T) {
		expected := "kirill"
		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1 {
			t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
		}
	})

	t.Run("test walk returns correct string", func(t *testing.T) {
		expected := "kirill"
		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if got[0] != expected {
			t.Errorf("wrong value returned from function, got %q want %q", got[0], expected)
		}
	})
}
