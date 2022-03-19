package structx

import (
	"errors"
	"reflect"
)

// Known errors
var (
	ErrNonStructType  = errors.New("type is not struct")
	ErrFieldNotExists = errors.New("field not exists in the struct")
	ErrNoMatchedField = errors.New("field that match given criteria does not exists")
)

// Introspect struct and return visible fields information or error.
// This function ignores unexported fields, channels and function members.
func Introspect[T any]() (*Info, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	i := Info{}
	if err := i.introspect(t); err != nil {
		return nil, err
	}
	return &i, nil

}

// MustIntrospect analyze struct T, return visible fields info, and panic if error.
// This function ignores unexported fields, channels and function members.
func MustIntrospect[T any]() *Info {
	t := reflect.TypeOf((*T)(nil)).Elem()
	i := Info{}
	if err := i.introspect(t); err != nil {
		panic(err)
	}
	return &i
}

// IntrospectFunc struct and return its visible fields related information
// filtered by supplied function fn.
func IntrospectFunc[T any](fn func(fi reflect.StructField) bool) (*Info, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	i := Info{}
	if err := i.introspectFunc(t, fn); err != nil {
		return nil, err
	}
	return &i, nil

}

// MustIntrospectFunc analyze struct T, and panic if error.
// Visible struct fields are filtered with supplied function fn.
func MustIntrospectFunc[T any](fn func(fi reflect.StructField) bool) *Info {
	t := reflect.TypeOf((*T)(nil)).Elem()
	i := Info{}
	if err := i.introspectFunc(t, fn); err != nil {
		panic(err)
	}
	return &i
}
