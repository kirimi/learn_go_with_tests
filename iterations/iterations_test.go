package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"
	assertCorrectRepetition(t, got, want)
}

func assertCorrectRepetition(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("Got %q Want %q", got, want)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Println(repeated)
	// output: aaaaaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
