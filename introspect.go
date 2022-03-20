package structx

import (
	"errors"
	"reflect"
)

// Known errors
var (
	ErrNonStructType  = errors.New("type is not struct")
	ErrFieldNotExists = errors.New("field not exists in the struct")
)

// Introspect struct and return visible fields information or error.
// This function ignores unexported fields, channels and function members.
func Introspect[T any]() (*Meta, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	m := Meta{}
	if err := m.introspect(t); err != nil {
		return nil, err
	}
	return &m, nil

}

// MustIntrospect analyze struct T, return visible fields info, and panic if error.
// This function ignores unexported fields, channels and function members.
func MustIntrospect[T any]() *Meta {
	t := reflect.TypeOf((*T)(nil)).Elem()
	m := Meta{}
	if err := m.introspect(t); err != nil {
		panic(err)
	}
	return &m
}

// IntrospectFunc struct and return its visible fields related information
// filtered by supplied function fn.
func IntrospectFunc[T any](fn func(fi reflect.StructField) bool) (*Meta, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	m := Meta{}
	if err := m.introspectFunc(t, fn); err != nil {
		return nil, err
	}
	return &m, nil

}

// MustIntrospectFunc analyze struct T, and panic if error.
// Visible struct fields are filtered with supplied function fn.
func MustIntrospectFunc[T any](fn func(fi reflect.StructField) bool) *Meta {
	t := reflect.TypeOf((*T)(nil)).Elem()
	m := Meta{}
	if err := m.introspectFunc(t, fn); err != nil {
		panic(err)
	}
	return &m
}

func (m *Meta) hasElement(ft reflect.Type) bool {
	return ft.Kind() == reflect.Pointer ||
		ft.Kind() == reflect.Slice ||
		ft.Kind() == reflect.Array
}

func (m *Meta) includeFunc(fi reflect.StructField) bool {
	if fi.Tag != "" {
		return true
	}

	ft := fi.Type
	for m.hasElement(ft) {
		ft = ft.Elem()
	}
	ki := ft.Kind()

	// Ignores fields that are:
	// 1. Unexported
	// 2. Embedded field
	// 3. Function member
	// 4. Channel
	ignore := !fi.IsExported() ||
		fi.Anonymous ||
		ki == reflect.Func ||
		ki == reflect.Chan

	return !ignore
}

// introspect reflect.Type of Struct
func (m *Meta) introspectStructFunc(t reflect.Type, fn func(f reflect.StructField) bool) error {
	fields := reflect.VisibleFields(t)
	for _, fi := range fields {
		if !fn(fi) {
			continue
		}

		tags, err := parseTag(fi.Tag)
		if err != nil {
			return err
		}
		sf := Field{
			SF:   fi,
			Name: fi.Name,
			Tags: tags,
			Zero: reflect.New(fi.Type),
		}

		m.Fields = append(m.Fields, sf)
	}
	return nil
}

func (m *Meta) introspect(t reflect.Type) error {
	return m.introspectFunc(t, m.includeFunc)
}

func (m *Meta) introspectFunc(t reflect.Type, fn func(f reflect.StructField) bool) error {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ErrNonStructType
	}
	m.Name = t.Name()
	m.PkgPath = t.PkgPath()

	return m.introspectStructFunc(t, fn)
}
