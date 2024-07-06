package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Raidus float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Raidus, 2)
}