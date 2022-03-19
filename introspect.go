package structx

import (
	"reflect"
)

func (in *Info) hasElement(ft reflect.Type) bool {
	return ft.Kind() == reflect.Pointer ||
		ft.Kind() == reflect.Slice ||
		ft.Kind() == reflect.Array
}

func (in *Info) includeFunc(fi reflect.StructField) bool {
	if fi.Tag != "" {
		return true
	}

	ft := fi.Type
	for in.hasElement(ft) {
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
func (in *Info) introspectStructFunc(t reflect.Type, fn func(f reflect.StructField) bool) error {
	fields := reflect.VisibleFields(t)
	for _, fi := range fields {
		if !fn(fi) {
			continue
		}

		// Get type or element's type
		ft := fi.Type
		ki := ft.Kind()
		for in.hasElement(ft) {
			ft = ft.Elem()
		}

		tags, err := parseTag(fi.Tag)
		if err != nil {
			return err
		}
		sf := Field{
			Kind:        ki,
			Name:        fi.Name,
			PkgPath:     fi.PkgPath,
			TypeName:    ft.Name(),
			TypePkgPath: ft.PkgPath(),
			Tags:        tags,
		}
		in.Fields = append(in.Fields, sf)
	}
	return nil
}

func (in *Info) introspect(t reflect.Type) error {
	return in.introspectFunc(t, in.includeFunc)
}

func (in *Info) introspectFunc(t reflect.Type, fn func(f reflect.StructField) bool) error {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ErrNonStructType
	}
	in.Name = t.Name()
	in.PkgPath = t.PkgPath()

	return in.introspectStructFunc(t, fn)
}
