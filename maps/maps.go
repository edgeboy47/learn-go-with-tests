package maps

import "errors"

type Dictionary map[string]string

var ErrValueNotFound = errors.New("Value Not Found")
var ErrKeyExists = errors.New("Key already in use")

func (dict Dictionary) Search(key string) (string, error) {
	val, ok := dict[key]

	if !ok {
		return "", ErrValueNotFound
	}

	return val, nil
}

func (dict Dictionary) Add(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrValueNotFound:
		dict[key] = value
		return nil
	case nil:
		return ErrKeyExists
	default:
		return err
	}
}
