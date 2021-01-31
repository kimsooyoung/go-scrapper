package mydict

import "errors"

// aliasing
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("Word Aleady Exists")
	errUpdateFail = errors.New("Update Fail cannot update non-existing Elements")
	errDeleteFail = errors.New("Deletion Fail cannot update non-existing Elements")
)

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

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errUpdateFail
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errDeleteFail
	}
	return nil
}
