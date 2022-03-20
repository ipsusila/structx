package structx

import (
	"reflect"
	"strings"
)

// Tag stores struct tag data
type Tag struct {
	Name    string
	Value   string
	Options []string
}

// Field stores struct field information
type Field struct {
	Name string
	Tags []Tag
	Zero reflect.Value
	SF   reflect.StructField
}

// Meta stores struct information
type Meta struct {
	Name    string
	PkgPath string
	Fields  []Field
}

// HasTag return true if tag with with given name,
// e.g. `json` exists.
func (f Field) HasTag(name string) bool {
	for _, t := range f.Tags {
		if t.Name == name {
			return true
		}
	}
	return false
}

// Tag return pointer to tag which name is given as arg.
func (f Field) Tag(name string) (*Tag, bool) {
	for _, t := range f.Tags {
		if t.Name == name {
			return &t, true
		}
	}
	return nil, false
}

// Identier return unique identifier for this struct.
// Identifier is composed by PkgPath + "." + Name [+ "." + key]
func (m Meta) Identifier(keys ...string) string {
	sb := strings.Builder{}
	if m.PkgPath != "" {
		sb.WriteString(m.PkgPath)
		sb.WriteByte('.')
	}
	sb.WriteString(m.Name)
	for _, key := range keys {
		sb.WriteByte('.')
		sb.WriteString(key)
	}
	return sb.String()
}

// String return PkgPath.Name for this struct
func (m Meta) String() string {
	return m.Identifier()
}
