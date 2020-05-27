package maps

import (
	"errors"
)

//ErrWordNotFound for not found error
var ErrWordNotFound = errors.New("Can't find the word")
var ErrWordExists = errors.New("Word is already exsists")

//Dictionary map
type Dictionary map[string]string

//Search in dictionary
func (d Dictionary) Search(needle string) (string, error) {
	defination, ok := d[needle]

	if !ok {
		return "", ErrWordNotFound
	}
	return defination, nil
}

//Add new word to dictionary
func (d Dictionary) Add(word string, defination string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = defination
	case nil:
		return ErrWordExists
	default:
		return err //unknown err
	}

	return nil
}

//Update existing word in the dictionary
func (d Dictionary) Update(word string, defination string) error {
	_, err := d.Search(word)

	if err == ErrWordNotFound {
		return err
	}

	d[word] = defination
	return nil
}

//Delete a word from a dictionary
func (d Dictionary) Delete(word string) error {
	delete(d, word)
	return nil
}
