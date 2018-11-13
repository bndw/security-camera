package main

import (
	"bytes"
	"fmt"
	"net/http"
)

//
func NotifySlack(webhookUrl, url string) error {
	// curl -X POST -H 'Content-type: application/json' --data '{"text":"https://s3...com/file.jpg"}' https://hooks.slack.com/services/some/webhook/path'

	var (
		msg     = []byte(fmt.Sprintf(`{"text":"%s"}`, url))
		payload = bytes.NewBuffer(msg)
	)

	req, err := http.NewRequest("POST", webhookUrl, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	return err
}
