package main

import (
	"context"
	"encoding/json"

	"github.com/ha-fr-av/hackathon-integration/lambdas/common"
)

type Handler struct {
}

type inputParams struct {
	InsuredID string `json:"insuredId"`
	PolicyID  string `json:"policyID"`
}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) Handle(ctx context.Context, event common.StepInput[inputParams]) (common.StepOutput[common.StepInput[common.QuoteResponseBody]], error) {
	output := common.StepOutput[common.StepInput[common.QuoteResponseBody]]{}

	a, err := arrange(event)
	if err != nil {
		msg := err.Error()
		output.Error = &msg
		return output, nil
	}

	res, err := act(a)
	if err != nil {
		msg := err.Error()
		output.Error = &msg
		return output, nil
	}

	var dat common.QuoteResponseBody
	if err := json.NewDecoder(res.Body).Decode(&dat); err != nil {
		msg := err.Error()
		output.Error = &msg
		return output, nil
	}

	if err = assert(input{StatusCode: res.StatusCode, Payload: dat}); err != nil {
		msg := err.Error()
		output.Error = &msg
		return output, nil
	}

	output.Payload.UserInfo = event.UserInfo
	output.Payload.Host = event.Host
	output.Payload.Payload = dat

	return output, nil
}
