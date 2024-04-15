package maps

import "testing"

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
  t.Helper()

  if got != want {
    t.Errorf("got error %q want %q", got, want)
  }
}

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "test string"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "test string"

		assertString(t, got, want)
	})

  t.Run("Unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := ErrValueNotFound

		assertError(t, err, want)
  })
}
