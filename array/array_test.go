package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of specified length", func(t *testing.T) {
		var actual int = 15
		numbers := [5]int{1, 2, 3, 4, 5}
    expected := Sum(numbers[:])

		if actual != expected {
			t.Errorf("got %d want %d, given %v", expected, actual, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
  t.Run("Multiple arrays of different length", func(t *testing.T) {
    arr1 := [] int {1, 2, 3}
    arr2 := [] int {4, 5}
    actual := SumAll(arr1, arr2)
    expected := [] int {6, 9}
    
    if !reflect.DeepEqual(actual, expected) {
      t.Errorf("got %v, want %v", actual, expected)
    }
  })

  t.Run("Handle empty arrays", func(t *testing.T) {
    arr1 := [] int {}
    arr2 := [] int {4}
    actual := SumAll(arr1, arr2)
    expected := [] int {0, 4}
    
    if !reflect.DeepEqual(actual, expected) {
      t.Errorf("got %v, want %v", actual, expected)
    }
  })
}
