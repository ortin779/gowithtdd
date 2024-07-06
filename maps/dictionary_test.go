package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		expected := "this is just a test"

		assertStrings(t, expected, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		expected := "could not find the word you were looking for"

		assertError(t, ErrNotFound, err)
		assertStrings(t, expected, err.Error())
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Add("test", "this is just a test")
		assertNoError(t, err)

		expected := "this is just a test"
		got, err := dict.Search("test")

		assertNoError(t, err)
		assertStrings(t, expected, got)
	})

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		err := dict.Add("test", "this is just a test")
		assertError(t, ErrAlreadyExist, err)

		expected := "this is just a test"
		got, err := dict.Search("test")

		assertNoError(t, err)
		assertStrings(t, expected, got)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		err := dict.Update("test", "new definition")
		assertNoError(t, err)

		expected := "new definition"
		got, err := dict.Search("test")

		assertNoError(t, err)
		assertStrings(t, expected, got)
	})

	t.Run("non existing word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Update("test", "this is just a test")
		assertError(t, ErrWordNotExist, err)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "test definition"}

	dict.Delete(word)

	_, err := dict.Search(word)

	assertError(t, ErrNotFound, err)
}

func assertStrings(t testing.TB, expected, got string) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}

func assertError(t testing.TB, expected, got error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected an error to be")
	}

	if expected != got {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatalf("didn't expect an error but received %q", got.Error())
	}
}
