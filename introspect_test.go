package structx_test

import (
	"net/http"
	"testing"

	"github.com/ipsusila/structx"
	"github.com/stretchr/testify/assert"
)

func TestIntrospect(t *testing.T) {
	i, err := structx.Introspect[http.Request]()
	assert.NoError(t, err, "Introspect shall not return error")
	printJson(i)

	i, err = structx.Introspect[int]()
	assert.Error(t, err, "Introspect int shall return error")

	i, err = structx.Introspect[myReq]()
	assert.NoError(t, err, "Introspect shall not return error")
	printJson(i)

	type myAlias myReq
	i, err = structx.Introspect[myAlias]()
	assert.NoError(t, err, "Introspect shall not return error")
	printJson(i)
}

var structName string

func BenchmarkIntrospect(b *testing.B) {
	var m *structx.Meta
	for i := 0; i < b.N; i++ {
		m, _ = structx.Introspect[myReq]()
	}
	structName = m.Name
}
