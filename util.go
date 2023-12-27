package FrostAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func ISO8601(timestamp string, daysToAdd, minutesToAdd int) string {
	parsedTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return ""
	}

	newTime := parsedTime
	if daysToAdd != 0 {
		newTime = newTime.AddDate(0, 0, daysToAdd)
	}

	if minutesToAdd != 0 {
		newTime = newTime.Add(time.Minute * time.Duration(minutesToAdd))
	}

	return newTime.Format(time.RFC3339)
}
