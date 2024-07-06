package maps

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var (
	ErrNotFound     = fmt.Errorf("could not find the word you were looking for")
	ErrAlreadyExist = errors.New("already exists a word in the dict")
	ErrWordNotExist = errors.New("can not update as word doesn't exist")
)

func (dict Dictionary) Search(word string) (string, error) {
	val, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (dict Dictionary) Add(word, meaning string) error {
	if _, ok := dict[word]; ok {
		return ErrAlreadyExist
	}
	dict[word] = meaning
	return nil
}

func (dict Dictionary) Update(word, definition string) error {
	_, err := dict.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordNotExist
	case nil:
		dict[word] = definition
	}
	return nil
}

func (dict Dictionary) Delete(word string) {
	delete(dict, word)
}
