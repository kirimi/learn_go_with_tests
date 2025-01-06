package main

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrKeyExist = errors.New("key exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, isExist := d[key]
	if isExist {
		return value, nil
	}

	return value, ErrNotFound
}

func (d Dictionary) Add(key, value string) error {
	if _, ok := d[key]; ok {
		return ErrKeyExist
	}

	d[key] = value

	return nil
}
