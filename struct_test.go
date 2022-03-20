package structx_test

import (
	"testing"
	"time"

	"github.com/ipsusila/structx"
	"github.com/stretchr/testify/assert"
)

func TestStructMapValues(t *testing.T) {
	d := myModel{
		ID:        100,
		FirstName: "John",
		CreatedAt: time.Now(),
	}
	dbFields := []string{"id", "first_name", "created_at", "updated_at"}
	dbVals := []interface{}{d.ID, d.FirstName, d.CreatedAt, d.UpdatedAt}
	dbPtrs := []interface{}{&d.ID, &d.FirstName, &d.CreatedAt, &d.UpdatedAt}

	s := structx.New[myModel]().
		MustIntrospect().
		MustMap("db")

	vals, err := s.MappedValues(&d, dbFields)
	assert.NoError(t, err, "MappedValues access shall not error")
	for i, v := range vals {
		assert.Equal(t, dbVals[i], v, "Value should be equal")
	}

	ptrs, err := s.MappedPointers(&d, dbFields)
	assert.NoError(t, err, "MappedPointers shall not error")
	for i, v := range ptrs {
		assert.Equal(t, dbPtrs[i], v, "Pointer should be equal")
	}
}
