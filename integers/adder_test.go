package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("should return sum of both numbers", func(t *testing.T) {
		sum := Add(5, 10)
		expected := 15

		compareInt(t, expected, sum)
	})

	t.Run("should return zero if both numbers zero", func(t *testing.T) {
		sum := Add(0, 0)
		expected := 0

		compareInt(t, expected, sum)
	})
}

func compareInt(t testing.TB, expected, received int) {
	t.Helper()

	if received != expected {
		t.Errorf("expected %d but got %d", expected, received)
	}
}

func ExampleAdd() {
	sum := Add(5, 7)
	fmt.Println(sum)

	//Output: 12
}
