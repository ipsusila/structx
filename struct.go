package structx

import (
	"fmt"
	"reflect"
)

// Struct access
type Struct[T any] struct {
	md *Meta
	sm StructMapper
}

// New return struct
func New[T any]() *Struct[T] {
	return &Struct[T]{}
}

// NewWithMeta return struct with given meta data.
// Optionally, mapper can be passed here
func NewWithMeta[T any](m *Meta, fm ...StructMapper) *Struct[T] {
	var sm StructMapper
	if len(fm) > 0 {
		sm = fm[0]
	}
	return &Struct[T]{md: m, sm: sm}
}

// WithMeta assign meta data and optionally StructMapper
func (a *Struct[T]) WithMeta(m *Meta, fm ...StructMapper) *Struct[T] {
	a.md = m
	if len(fm) > 0 {
		a.sm = fm[0]
	}
	return a
}

// Metadata return struct meta data
func (a *Struct[T]) Metadata() *Meta {
	return a.md
}

// WithMetadata assign metadata
func (a *Struct[T]) WithMetadata(m *Meta) *Struct[T] {
	a.md = m
	return a
}

// Mapper return struct mapper
func (a *Struct[T]) Mapper() StructMapper {
	return a.sm
}

// WithMapper assign field mapper
func (a *Struct[T]) WithMapper(fm StructMapper) *Struct[T] {
	a.sm = fm
	return a
}

// Introspect struct meta data
func (a *Struct[T]) Introspect() error {
	m, err := Introspect[T]()
	if err != nil {
		return err
	}
	a.md = m
	return nil
}

// MustIntrospect introspect struct meta data and panic if error.
func (a *Struct[T]) MustIntrospect() *Struct[T] {
	if err := a.Introspect(); err != nil {
		panic(err)
	}
	return a
}

// Map maps struct based on tag or it's value
func (a *Struct[T]) Map(tag ...string) error {
	fm, err := MapStruct[T](tag...)
	if err != nil {
		return err
	}
	a.sm = fm
	return nil
}

// MustMap struct fields
func (a *Struct[T]) MustMap(tag ...string) *Struct[T] {
	if err := a.Map(tag...); err != nil {
		panic(err)
	}
	return a
}

// FieldValues extract field values from given v
func (a *Struct[T]) FieldValues(v *T) []interface{} {
	fields := a.md.Fields
	values := make([]interface{}, len(fields))
	rv := reflect.ValueOf(v).Elem()
	for i := 0; i < len(fields); i++ {
		idx := fields[i].SF.Index
		fv := rv.FieldByIndex(idx)
		values[i] = fv.Interface()
	}

	return values
}

// FieldPointers extract field address (pointer) from given v
func (a *Struct[T]) FieldPointers(v *T) []interface{} {
	fields := a.md.Fields
	values := make([]interface{}, len(fields))
	rv := reflect.ValueOf(v).Elem()
	for i := 0; i < len(fields); i++ {
		idx := fields[i].SF.Index
		fv := rv.FieldByIndex(idx)
		values[i] = fv.Addr().Interface()
	}

	return values
}

// MappedValues extract field values from given v using mapper
func (a *Struct[T]) MappedValues(v *T, keys []string) ([]interface{}, error) {
	values := make([]interface{}, len(keys))
	rv := reflect.ValueOf(v).Elem()
	for i, key := range keys {
		idx, ok := a.sm.Index(key)
		if !ok {
			return nil, fmt.Errorf("field %s does not exists", key)
		}
		fv := rv.FieldByIndex(idx)
		values[i] = fv.Interface()
	}

	return values, nil
}

// MappedPointers extract field values from given v using mapper
func (a *Struct[T]) MappedPointers(v *T, keys []string) ([]interface{}, error) {
	values := make([]interface{}, len(keys))
	rv := reflect.ValueOf(v).Elem()
	for i, key := range keys {
		idx, ok := a.sm.Index(key)
		if !ok {
			return nil, fmt.Errorf("field %s does not exists", key)
		}
		fv := rv.FieldByIndex(idx)
		values[i] = fv.Addr().Interface()
	}

	return values, nil
}
