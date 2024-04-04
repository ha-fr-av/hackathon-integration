package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ha-fr-av/hackathon-integration/lambdas/common"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) Handle(ctx context.Context, event map[string]any) (common.StepResponse, error) {
	var output common.StepResponse

	a, err := arrange(event)
	if err != nil {
		return output, err
	}

	res, err := act(a)
	if err != nil {
		return output, err
	}

	var dat QuoteResponseBody
	if err := json.NewDecoder(res.Body).Decode(&dat); err != nil {
		return output, err
	}

	if res.StatusCode != http.StatusCreated {
		return output, errors.New("Invalid response status code")
	}

	if err = assert(dat); err != nil {
		return output, err
	}

	output.Payload = dat

	return output, nil
}
