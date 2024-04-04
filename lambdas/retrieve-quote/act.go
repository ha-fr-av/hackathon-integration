package main

import (
	"net/http"
)

type ActionParams struct {
	Endpoint string
	Headers  map[string]string
	QuoteID  string
}

func act(a ActionParams) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, a.Endpoint, nil)
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
