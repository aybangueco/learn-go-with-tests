package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat a 30 times", func(t *testing.T) {
		repeated := Repeat("a", 30)
		expected := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("replace character with 6 x", func(t *testing.T) {
		repeated := strings.ReplaceAll(Repeat("a", 6), "aaaaaa", "xxxxxx")
		expected := "xxxxxx"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("capitalize a repeated character", func(t *testing.T) {
		repeated := strings.ToUpper(Repeat("a", 5))
		expected := "AAAAA"

		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got repeated %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
