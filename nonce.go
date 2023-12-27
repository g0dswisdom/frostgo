package FrostAPI

import (
	"fmt"
	"time"
)

var epoch = time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6
var increment int64 = 0

func NewNonce() string {
	timestamp := time.Now().UnixNano() / 1e6

	if increment >= 4095 {
		increment = 0
	}

	if timestamp < epoch {
		panic("Invalid timestamp")
	}

	result := ((timestamp - epoch) << 22) | (1 << 17) | increment
	increment++

	return fmt.Sprintf("%d", result)
}
