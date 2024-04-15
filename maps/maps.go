package maps

import "errors"

type Dictionary map[string]string

var ErrValueNotFound = errors.New("Value Not Found")

func (dict Dictionary) Search(key string) (string, error) {
	val, ok := dict[key]

	if !ok {
		return "", ErrValueNotFound
	}

	return val, nil
}
