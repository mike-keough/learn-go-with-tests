package structs

import "testing"

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0, 10.0}
// 	got := Perimeter(rectangle)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("Got: %.2f Want: %.2f", got, want)
// 	}
// }

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{w: 12, h: 6}, want: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%v got %g want %g", tt.shape, got, tt.want)
			}
		})

	}

}
