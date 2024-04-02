package array

import "testing"

func TestArray(t *testing.T) {
	t.Run("Collection of specified length", func(t *testing.T) {
		var actual int = 15
		numbers := [5]int{1, 2, 3, 4, 5}
    expected := Sum(numbers[:])

		if actual != expected {
			t.Errorf("got %d want %d, given %v", expected, actual, numbers)
		}
	})
}
