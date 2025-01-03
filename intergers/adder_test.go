package intergers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Integers", func(t *testing.T) {
		sum := Add(2, 2)
		want := 4
		assertCorrectSum(t, sum, want)
	})
}

func assertCorrectSum(t testing.TB, want, got int) {
	if got != want {
		t.Errorf("Got %d - Want %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Println(sum)
	// Output: 7
}
