package FrostAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Helper function to make requests. Sometimes refuses to work with responses that return arrays.
//
// Deprecated: I don't really like this function
func customRequest(b *Bot, method, endpoint string, data map[string]interface{}, headers map[string]interface{}) *http.Response {
	response, err := b.Request(true, method, endpoint, data, nil)
	if err != nil {
		fmt.Printf("Could not make %s request: %s", method, err.Error())
		return nil
	}
	return response
}

// Helper function for decoding JSON.
func decode(resp *http.Response, target interface{}) error {
	err := json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

var epoch = time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6
var increment int64 = 0

func newNonce() string {
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
