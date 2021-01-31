package mydict

import "errors"

// aliasing
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("Word Aleady Exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	} else {
		return "", errNotFound
	}
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
