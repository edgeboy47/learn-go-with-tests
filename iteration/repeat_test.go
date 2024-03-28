package iteration

import "testing"

func TestRepeat(t *testing.T) {
  actual := repeat("a", 5)
  expected := "aaaaa"

  if actual != expected {
    t.Errorf("expected %q but got %q", expected, actual)
  }
}
