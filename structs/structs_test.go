package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	actual := Perimeter(rectangle)
	expected := 40.0

	if actual != expected {
		t.Errorf("got %.2f want %.2f", actual, expected)
	}
}

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.0, 4.0}

		actual := rectangle.Area()
		expected := 20.0

		if actual != expected {
			t.Errorf("got %.2f want %.2f", actual, expected)
		}
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.0}

		actual := circle.Area()
		expected := 314.1592653589793

		if actual != expected {
			t.Errorf("got %.2f want %.2f", actual, expected)
		}
	})
}
