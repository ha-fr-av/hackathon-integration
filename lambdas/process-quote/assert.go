package main

import (
	"errors"
	"net/http"

	"github.com/ha-fr-av/hackathon-integration/lambdas/common"
)

type input struct {
	Payload    common.QuoteResponseBody
	StatusCode int
}

func assert(dat input) error {

	if dat.StatusCode != http.StatusCreated {
		return errors.New("Invalid response status code")
	}

	if dat.Payload.QuoteID == "" {
		return errors.New("QuoteId not returned")
	}

	if dat.Payload.QuoteRejected {
		return errors.New("Quote Rejected")
	}

	return nil

}
