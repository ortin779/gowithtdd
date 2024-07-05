package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat('c')
	expected := "ccccc"

	if got != expected {
		t.Errorf("got %q but expected %q", got, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat('a')
	}
}
