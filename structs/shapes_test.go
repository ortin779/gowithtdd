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
		name     string
		shape    Shape
		expected float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{
				Height: 12.0,
				Width:  6.0,
			},
			expected: 72.0,
		},
		{
			name: "Circle",
			shape: Circle{
				Raidus: 10.0,
			},
			expected: 314.1592653589793,
		},
		{
			name: "Triangle",
			shape: Triangle{
				Base:   12,
				Height: 6,
			},
			expected: 36,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.expected {
				t.Errorf("%#v expected %g, but got %g", tt.shape, tt.expected, got)
			}
		})
	}

}
