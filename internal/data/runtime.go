package data

import (
	"fmt"
	"strconv"
)

// Runtime represents the movie length
type Runtime int32

// MarshalJSON implements json.Marshler interface to return custom json for Runtime type
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
