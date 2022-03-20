package structx_test

import (
	"testing"

	"github.com/ipsusila/structx"
	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {
	tags, err := structx.Tags[myReq]("ABC")
	assert.NoError(t, err, "Tags ABC shall not return error")
	printJson(tags)

	tags, err = structx.TagsByName[myReq]("json")
	assert.NoError(t, err, "Extract `json` tag shall not return error")
	printJson(tags)
}
