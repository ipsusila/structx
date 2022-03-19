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
	Kind        reflect.Kind
	Name        string
	PkgPath     string
	TypeName    string
	TypePkgPath string
	Tags        []Tag
}

// Info stores struct information
type Info struct {
	Name    string
	PkgPath string
	Fields  []Field
}

// Identier return unique identifier for this struct.
// Identifier is composed by PkgPath + "." + Name [+ "." + key]
func (i Info) Identifier(keys ...string) string {
	sb := strings.Builder{}
	if i.PkgPath != "" {
		sb.WriteString(i.PkgPath)
		sb.WriteByte('.')
	}
	sb.WriteString(i.Name)
	for _, key := range keys {
		sb.WriteByte('.')
		sb.WriteString(key)
	}
	return sb.String()
}

// String return PkgPath.Name for this struct
func (i Info) String() string {
	return i.Identifier()
}
