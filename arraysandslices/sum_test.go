package arraysandslices

import (
	"reflect"
	"slices"
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

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, expected, got []int) {
		t.Helper()

		if !slices.Equal(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{1, 2, 3})
		expect := []int{2, 5}

		checkSums(t, expect, got)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, want, got)

	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([][]int{{1, 2, 3}, {2, 3}}...)
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([][]int{{1, 2, 3}, {2, 3}}...)
	}
}
