package structx_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ipsusila/structx"
)

type myReq struct {
	fmt.Stringer
	*http.Request
	ID     string       `json:"id"`
	Origin string       `json:"origin"`
	Info   structx.Meta `json:"info"`

	ABC struct {
		ID string
	} `json:"abc,omitempty" db:"abc,soft_delete"`

	Data   interface{}
	Fields []*structx.Field
	_      string `table:"my_request"`
}

type myModel struct {
	ID        int        `json:"id" db:"id"`
	FirstName string     `json:"firstName" db:"first_name"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at,createAt"`
	UpdatedAt *time.Time `db:"updated_at,updatedAt"`
}

func printJson(v interface{}) {
	js, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(js))
}
