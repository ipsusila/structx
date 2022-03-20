package structx_test

import (
	"testing"

	"github.com/ipsusila/structx"
	"github.com/stretchr/testify/assert"
)

func TestMapper(t *testing.T) {
	const (
		FirstName  = "FirstName"
		firstName  = "firstName"
		first_name = "first_name"
	)
	// Map by struct name
	m, err := structx.MapStruct[myModel]()
	assert.NoError(t, err, "Map by field shall not error")
	fi, ok := m.Field(FirstName)
	assert.True(t, ok, FirstName+" should exist in mapper")
	assert.Equal(t, FirstName, fi.Name, "Field name should be "+FirstName)

	// Map by "json"
	m, err = structx.MapStruct[myModel]("json")
	assert.NoError(t, err, "Map by `json` tag shall not error")
	fi, ok = m.Field(firstName)
	assert.True(t, ok, firstName+" should exist in mapper")
	assert.Equal(t, FirstName, fi.Name, "Field name should be "+FirstName)
	fi, ok = m.Field("updateAt")
	assert.False(t, ok, "Field updateAt shall not exists")

	// Map by "db" tag
	m, err = structx.MapStruct[myModel]("db")
	assert.NoError(t, err, "Map by `db` tag shall not error")
	fi, ok = m.Field(first_name)
	assert.True(t, ok, first_name+" should exist in mapper")
	assert.Equal(t, FirstName, fi.Name, "Field name should be "+FirstName)
}
