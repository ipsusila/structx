package structx_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ipsusila/structx"
	"github.com/stretchr/testify/assert"
)

type myReq struct {
	fmt.Stringer
	*http.Request
	ID     string       `json:"id"`
	Origin string       `json:"origin"`
	Info   structx.Info `json:"info"`

	ABC struct {
		ID string
	} `json:"abc,omitempty" db:"abc,soft_delete"`

	Data   interface{}
	Fields []*structx.Field
	_      string `table:"my_request"`
}

func printJson(v interface{}) {
	js, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(js))
}

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

func TestTags(t *testing.T) {
	tags, err := structx.Tags[myReq]("ABC")
	assert.NoError(t, err, "Tags ABC shall not return error")
	printJson(tags)

	tags, err = structx.TagsByName[myReq]("json")
	assert.NoError(t, err, "Extract `json` tag shall not return error")
	printJson(tags)
}
