package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{
		Height: 10.0,
		Width:  10.0,
	}
	got := Perimeter(rect)
	expected := 40.0

	if got != expected {
		t.Errorf("expected %.2f, but got %.2f", expected, got)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()

		got := shape.Area()
		if expected != got {
			t.Errorf("expected %.2f, but got %.2f", expected, got)
		}
	}
	t.Run("rectangles area", func(t *testing.T) {
		rect := Rectangle{
			Height: 12.0,
			Width:  6.0,
		}
		checkArea(t, rect, 72.0)
	})

	t.Run("circle area", func(t *testing.T) {
		c := Circle{Raidus: 10.0}
		checkArea(t, c, 314.1592653589793)

	})
}
