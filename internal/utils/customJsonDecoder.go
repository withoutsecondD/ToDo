package utils

import (
	"bytes"
	"encoding/json"
)

// CustomJSONDecoder is used to disallow request bodies with unknown fields,
// i.e. ones that are not listed in struct fields
func CustomJSONDecoder(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}
