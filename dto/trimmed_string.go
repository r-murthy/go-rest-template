package model

import (
	"encoding/json"
	"strings"
)

// TrimmedString is a string type without leading and trailing white spaces in the underlying string.
type TrimmedString string

// UnmarshalJSON deserializes vanilla strings and removes leading and trailing white spaces from them.
func (ts *TrimmedString) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return
	}
	*ts = TrimmedString(strings.TrimSpace(s))
	return
}
