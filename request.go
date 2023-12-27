package FrostAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (b *Bot) Request(auth bool, method, endpoint string, data interface{}, headers map[string]interface{}) (*http.Response, error) {
	autoHeaders := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9018 Chrome/108.0.5359.215 Electron/22.3.24 Safari/537.36",
	}

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}
	if auth {
		req.Header.Set("Authorization", b.Token)
	}
	for key, value := range autoHeaders {
		req.Header.Set(key, value)
	}

	for key, value := range headers {
		req.Header.Set(key, fmt.Sprintf("%v", value))
	}

	switch v := data.(type) {
	case map[string]interface{}:
		req.Header.Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		req.Body = ioutil.NopCloser(strings.NewReader(string(jsonData)))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
