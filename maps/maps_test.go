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

func assertValue(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, definition)
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

func TestAdd(t *testing.T) {
	dict := Dictionary{}

	t.Run("New Key", func(t *testing.T) {
		err := dict.Add("key", "value")
		key := "key"
		want := "value"

		assertError(t, err, nil)
		assertValue(t, dict, key, want)
	})

	t.Run("Existing Key", func(t *testing.T) {
		err := dict.Add("key", "value2")
		key := "key"
		want := "value"

		assertError(t, err, ErrKeyExists)
		assertValue(t, dict, key, want)
	})
}
