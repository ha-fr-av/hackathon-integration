package main

import (
	"encoding/json"
	"errors"
	"net/http"
)


type QuoteResponseBody struct {
	QuoteID       string `json:"quoteId"`
	QuoteRejected bool   `json:"quoteRejected"`
}

func assert(h *http.Response) error {

	if h.StatusCode != http.StatusCreated {
		return errors.New("Invalid response status code")
	}

	var dat QuoteResponseBody

	err := json.NewDecoder(h.Body).Decode(&dat)

	if err != nil {
		return errors.New("Cannot parse response body")
	}

	if dat.QuoteID == "" {
		return errors.New("QuoteId not returned")
	}

	if dat.QuoteRejected {
		return errors.New("Quote Rejected")
	}

	return nil

}
