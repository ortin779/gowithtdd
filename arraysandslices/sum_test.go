package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d but expected %d given, %v", got,
				expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	collections := [][]int{
		{1, 2, 3},
		{1, 2},
	}
	got := SumAll(collections...)
	expected := []int{6, 3}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([][]int{{1, 2, 3}, {2, 3}}...)
	}
}
