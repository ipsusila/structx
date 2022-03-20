package structx

import "reflect"

// StructMapper handles mapping struct field.
type StructMapper interface {
	StructMap(tag string, m *Meta)
	Field(key string) (Field, bool)
	Index(key string) ([]int, bool)
}

// MapStruct create field mapping for given struct T based on tag.
// If tag is empty (""), each field will be mapped by it's name.
func MapStruct[T any](tag ...string) (StructMapper, error) {
	var mtag string
	if len(tag) > 0 {
		mtag = tag[0]
	}
	m, err := Introspect[T]()
	if err != nil {
		return nil, err
	}
	sm := newStructMapper()
	sm.StructMap(mtag, m)

	return sm, nil
}

// MapStructFunc create field mapping for given struct T
// based on tag and filter function fn.
// If tag is empty (""), each field will be mapped by it's name.
func MapStructFunc[T any](fn func(fi reflect.StructField) bool, tag ...string) (StructMapper, error) {
	var mtag string
	if len(tag) > 0 {
		mtag = tag[0]
	}
	m, err := IntrospectFunc[T](fn)
	if err != nil {
		return nil, err
	}
	sm := newStructMapper()
	sm.StructMap(mtag, m)

	return sm, nil
}

type structMapper struct {
	fm map[string]Field
	im map[string][]int
}

func newStructMapper() *structMapper {
	return &structMapper{
		fm: make(map[string]Field),
		im: make(map[string][]int),
	}
}

func (s *structMapper) mapByTag(name string, m *Meta) {
	for _, fi := range m.Fields {
		if tag, ok := fi.Tag(name); ok {
			s.fm[tag.Value] = fi
			s.im[tag.Value] = fi.SF.Index
		} else {
			s.fm[fi.Name] = fi
			s.im[fi.Name] = fi.SF.Index
		}
	}
}
func (s *structMapper) mapByName(m *Meta) {
	for _, fi := range m.Fields {
		s.fm[fi.Name] = fi
		s.im[fi.Name] = fi.SF.Index
	}
}

func (s *structMapper) StructMap(tag string, m *Meta) {
	if tag == "" {
		s.mapByName(m)
	} else {
		s.mapByTag(tag, m)
	}
}
func (t *structMapper) Field(key string) (Field, bool) {
	fi, ok := t.fm[key]
	return fi, ok
}
func (t *structMapper) Index(key string) ([]int, bool) {
	i, ok := t.im[key]
	return i, ok
}
