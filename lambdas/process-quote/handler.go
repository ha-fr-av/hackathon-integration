package main

import (
	"context"

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

	if err = assert(res); err != nil {
		return output, err
	}

	output.Payload = "OK"
	output.Status = 200

	return output, nil
}
