package structx

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// Known errors
var (
	ErrNoMatchedField = errors.New("field that match given criteria does not exists")
)

func emptyOrContains(items []string, value string) bool {
	if len(items) == 0 {
		return true
	}
	for _, v := range items {
		if v == value {
			return true
		}
	}
	return false
}

func parseTag(tag reflect.StructTag, namesFilter ...string) ([]Tag, error) {
	tags := []Tag{}
	for tag != "" {
		// Skip leading space.
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		// note: ips, code taken from go reflection with some modification.
		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := string(tag[:i])
		tag = tag[i+1:]

		// Scan quoted string to find value.
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		qvalue := string(tag[:i+1])
		tag = tag[i+1:]

		// Extract Tag name, value and options
		if emptyOrContains(namesFilter, name) {
			value, err := strconv.Unquote(qvalue)
			if err != nil {
				return tags, err
			}
			items := strings.Split(value, ",")
			tv := Tag{
				Name: name,
			}
			if len(items) > 0 {
				tv.Value = items[0]
			}
			if len(items) > 1 {
				tv.Options = items[1:]
			}
			tags = append(tags, tv)
		}
	}
	return tags, nil
}

// Tags return parsed tags for given field name
func Tags[T any](fieldName string) ([]Tag, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	if t.Kind() != reflect.Struct {
		return nil, ErrNonStructType
	}
	fi, ok := t.FieldByName(fieldName)
	if !ok {
		return nil, ErrFieldNotExists
	}
	return parseTag(fi.Tag)
}

// TagsFunc return parsed tags.
// Only first field that satisfy given function will be returned.
func TagsFunc[T any](fn func(fi reflect.StructField) bool) ([]Tag, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	if t.Kind() != reflect.Struct {
		return nil, ErrNonStructType
	}

	tags := []Tag{}
	fields := reflect.VisibleFields(t)
	for _, fi := range fields {
		if fn(fi) {
			fiTags, err := parseTag(fi.Tag)
			if err != nil {
				return tags, err
			}
			tags = append(tags, fiTags...)
		}
	}
	if len(tags) == 0 {
		return nil, ErrNoMatchedField
	}
	return tags, nil
}

// TagsByName return list of tags for given tag name
func TagsByName[T any](name string) ([]Tag, error) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	if t.Kind() != reflect.Struct {
		return nil, ErrNonStructType
	}
	tags := []Tag{}
	fields := reflect.VisibleFields(t)
	for _, fi := range fields {
		fiTags, err := parseTag(fi.Tag, name)
		if err != nil {
			return tags, err
		}
		tags = append(tags, fiTags...)
	}
	return tags, nil
}
