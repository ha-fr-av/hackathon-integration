package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ActionParams struct {
	Endpoint string            `json:"Host"`
	Payload  map[string]any    `json:"payload"`
	Headers  map[string]string `json:"headers"`
}

func act(a ActionParams) (*http.Response, error) {
	j, _ := json.Marshal(a.Payload)

	req, err := http.NewRequest(http.MethodPost, a.Endpoint, bytes.NewReader(j))
	if err != nil {
		return nil, err
	}

	for key, val := range a.Headers {
		req.Header.Set(key, val)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	return resp, err
}
