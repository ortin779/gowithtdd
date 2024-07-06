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

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{
			Rectangle{
				Height: 12.0,
				Width:  6.0,
			},
			72.0,
		},
		{
			Circle{
				Raidus: 10.0,
			},
			314.1592653589793,
		},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.expected {
			t.Errorf("expected %g, but got %g", tt.expected, got)
		}
	}

}
