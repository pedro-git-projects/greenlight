package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf("%d mins", r)
	json = strconv.Quote(json)
	return []byte(json), nil
}
