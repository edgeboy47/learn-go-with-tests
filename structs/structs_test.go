package structs

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{5.0, 4.0}, 18.0},
		{Circle{10.0}, 62.83185307179586},
	}

	for _, shape := range perimeterTests {
		actual := shape.shape.Perimeter()
		if actual != shape.expected {
			t.Errorf("got %g want %g", actual, shape.expected)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name:     "Rectangle",
			shape:    Rectangle{5.0, 4.0},
			expected: 20.0,
		},
		{
			name:     "Circle",
			shape:    Circle{10.0},
			expected: 314.1592653589793,
		},
		{
			name:     "Triangle",
			shape:    Triangle{10.0, 5.0},
			expected: 25.0,
		},
	}

	for _, shape := range areaTests {
		t.Run(shape.name, func(t *testing.T) {
			actual := shape.shape.Area()
			if actual != shape.expected {
				t.Errorf("%v got %g want %g", shape.shape, actual, shape.expected)
			}
		})
	}
}
