package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aviva-verde/policy"
)

type event struct{}

var headers = map[string]string{
	"policyholder-dob": "1989-03-23",
}

// hardcoded items for initial dev
const quoteId = "692821c5-9b0a-4eb0-aaad-e355ee1c2fe9"
const domain = "https://m3yktt0wkj.execute-api.eu-west-1.amazonaws.com"
const jwtToken = "<<FOR DEV PUT JWT TOKEN HERE>>"

func Handler(ctx context.Context, event event) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, getEndpoint(quoteId), nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	req.Header.Set("policyholder-dob", "1989-03-23")
	req.Header.Set("Host", "localhost")

	client := &http.Client{}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	validationErr := validateTest(resp)

	if validationErr != nil {
		return resp, validationErr
	}

	return resp, nil

}

/**
* build the endpoint to get an mta quote
**/
func getEndpoint(quoteId string) string {

	//  "https://m3yktt0wkj.execute-api.eu-west-1.amazonaws.com/prod/quote/:quoteId"

	return fmt.Sprintf("%s/prod/quote/%s", domain, quoteId)

}

func main() {

	data, err := Handler(context.Background(), event{})
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("%v", data)
}

func validateTest(h *http.Response) error {

	if h.StatusCode != http.StatusOK {
		return errors.New("Invalid response status code")
	}

	var quote policy.Policy

	err := json.NewDecoder(h.Body).Decode(&quote)

	if err != nil {
		return errors.New("Cannot parse response body")
	}

	if quote.Adjustments[0].Type != "VehicleDetailsChanged" {
		return errors.New("Invalid Quote Adjustments")
	}

	return nil

}
