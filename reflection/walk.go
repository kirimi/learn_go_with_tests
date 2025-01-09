package main

import "reflect"

// see https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reflection
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
