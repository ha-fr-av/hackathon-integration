package main

import (
	"context"
	"encoding/json"

	"github.com/ha-fr-av/hackathon-integration/lambdas/common"
)

type Handler struct {
}

type inputParams struct {
	InsuredID string `json:"insurerId"`
	PolicyID  string `json:"policyID"`
}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) Handle(ctx context.Context, event common.StepInput[inputParams]) (common.StepOutput[common.StepInput[common.QuoteResponseBody]], error) {
	output := common.StepOutput[common.StepInput[common.QuoteResponseBody]]{}
	output.Payload.UserInfo = event.UserInfo
	output.Payload.Host = event.Host

	a, err := arrange(event)
	if err != nil {
		output.Error = err.Error()
		return output, nil
	}

	res, err := act(a)
	if err != nil {
		output.Error = err.Error()
		return output, nil
	}

	var dat common.QuoteResponseBody
	if err := json.NewDecoder(res.Body).Decode(&dat); err != nil {
		output.Error = err.Error()
		return output, nil
	}

	if err = assert(input{StatusCode: res.StatusCode, Payload: dat}); err != nil {
		output.Error = err.Error()
		return output, nil
	}

	output.Payload.Payload = dat

	return output, nil
}
