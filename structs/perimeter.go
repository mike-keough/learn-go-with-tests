package structs

import "math"

type Rectangle struct {
	w float64
	h float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func (r Rectangle) Area() float64 {
	return r.h * r.w
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
