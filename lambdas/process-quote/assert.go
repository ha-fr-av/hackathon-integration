package main

import (
	"errors"
)

type QuoteResponseBody struct {
	QuoteID       string `json:"quoteId"`
	QuoteRejected bool   `json:"quoteRejected"`
}

func assert(dat QuoteResponseBody) error {
	if dat.QuoteID == "" {
		return errors.New("QuoteId not returned")
	}

	if dat.QuoteRejected {
		return errors.New("Quote Rejected")
	}

	return nil

}
