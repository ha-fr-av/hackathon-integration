package main

import (
	"errors"
	"net/http"

	"github.com/aviva-verde/policy"
)

type input struct {
	Payload    policy.Policy
	StatusCode int
}

func assert(dat input) error {

	if dat.StatusCode != http.StatusOK {
		return errors.New("Invalid response status code")
	}

	if dat.Payload.Adjustments[0].Type != "VehicleDetailsChanged" {
		return errors.New("Invalid Quote Adjustments")
	}

	return nil

}
