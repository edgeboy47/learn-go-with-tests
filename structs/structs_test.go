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
		shape    Shape
		expected float64
	}{
		{Rectangle{5.0, 4.0}, 20.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, shape := range areaTests {
		actual := shape.shape.Area()
		if actual != shape.expected {
			t.Errorf("got %g want %g", actual, shape.expected)
		}
	}
}
