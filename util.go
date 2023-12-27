package FrostAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func customRequest(b *Bot, method, endpoint string, data map[string]interface{}, headers map[string]interface{}) *http.Response {
	response, err := b.Request(true, method, endpoint, data, nil)
	if err != nil {
		fmt.Printf("Could not make %s request: %s", method, err.Error())
		return nil
	}
	return response
}

func decode(resp *http.Response, target interface{}) {
	json.NewDecoder(resp.Body).Decode(target)
}
