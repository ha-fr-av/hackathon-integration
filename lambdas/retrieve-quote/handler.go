package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aviva-verde/policy"
	"github.com/ha-fr-av/hackathon-integration/lambdas/common"
)

type QuoteResponseBody struct {
	QuoteID       string `json:"quoteId"`
	QuoteRejected bool   `json:"quoteRejected"`
}

type event struct {
	Data QuoteResponseBody `json:"data"`
}

var headers = map[string]string{
	"policyholder-dob": "1989-03-23",
}

type Handler struct{}

// hardcoded items for initial dev
// const quoteId = "692821c5-9b0a-4eb0-aaad-e355ee1c2fe9"
const domain = "https://m3yktt0wkj.execute-api.eu-west-1.amazonaws.com"
const jwtToken = "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ6ZXJvLXRlc3RpbmcuYXZpdmEuY28udWsiLCJleHAiOjE3MTIyNjQwOTQsImlhdCI6MTcxMjIyMDg5NCwibmJmIjoxNzEyMjIwODk0LCJzdWIiOiJAZW5jOkt3NFJNSmtlRDRCbXNBazJzZ2txamZseDI5bXlCYXUrSk8zbDNYMFd6YStTSjgwS21RVnhLbktnd2xPcVBtaDltWEZMNE9rRkxTR2ZxUVdnNU5KR2dLYWlJYTV3eU5OYWEvbHdNZDBoYXNvRkxrUWdsT2F5TlhDOWZlaDM2c1I0V0R2cTdxMlN2clVHTWZtQVJrZnY3VnJITk0xWUdnS3ZTMXphblBWU3FpOUZJSVNhVjhxM05WWEZYYkhZcmVrRzA4bGJDaXdFcnQ3d1BYd1dObWFRSGFqOFlhWlVmTHhyN2JreTQzT0x2UllZNTEza0VIK2NJRnQ5MzAvUUFQajNuc1FTM0p3aG5qQ2QrenUzQ21HbnNZTjhwVDlaVTZwTDJSTUZycVp5Q2NXeXJMR09xd0pwQWtsb0JpajU3MmNKMmV2MFRjU1JBdHNaVzVkbFQwNXdFVkRwQ29hZ0I1THFBZ0RLY0NneWx6cDdpQ0RJZUlpNmRnK3pVdG5yYkhNSUtYZVhqM05ieWhTcDlYVDRLc2VsVEpMWlNrVC8yUnl6dkNDSEx3RGQwTktGMkpHaS9FalhUNFBtSkJqWFJiK1FJNi9aQ2M1NnEyM011QWZpOS9HMEZ4MUpLUlU4T0lkTEJpNWhQOS9jaXp6M1pIVVhEbHRTK0MxVGVnZzBDRjhjaVZWanY1U2xwZ051RVBKeEFBcTVKNHlyd3lSb2N1WWVPVlVIbm95TENyTVF2SHdRRzRMSGFjSW1TYW9zVmM2WFM2dFl3azB0cHdYeGIrWk9aMEQ5NERHUTNBY1M2MzVackp5enZ2VDlRWVZPK3NEVmN2Z09xdUxyY01mQjFVL0JsWjRabXZPbUY2amRHd2dTY1BnUHp1MDE0SmFQOVBCajlJOUNDd1I1TVEwPSJ9.wLBE-DLUCQOhAwzTbw9HSAoKQg2G9YuZDk1sLd0X0TCSKxKwUiM31JEBv4M-Y7t-NKDcxy2qUJqDvze3T8YlBXGBgdREj3uSLqOCfFHwPdjvjYopletap42hfUskcUg-vHm7AGBYrTOWFTFZ80mstwj0O9LqEZ0Qj_9gndwURVOBmbaL-9R0wcr-Vjh_2PGf-ml9d_0DNXYyWnT8IN884wEcax5hjZ4g8olZ___9JIv8IcHOho591sktKVIOuSW05RzZTD2yjs40inw6hZi7h9ICjZI-WCwOPVXmtz9VLAyVHhI--vJffOn2BhjnfCvNcnmZ7TevdwPrKeARGh-Ae6K9oTAyRpzVenvcaqDszd5_eCqD6XpXFZOqS_HtcJB5XKrGvpRf-c7gK0bY5U6C1OoTdWkiGw4ddDFDGtZjJ2027tP-hb_RLSa0fndfhfsK9vOw_CyM1Q2nGPQ3u00bNohglQPrPSDvRfu4BNZhI8y3Q1kAkBkutwkpK3kHQypAqsa_YUlZQGoqnXE-_v6uVDnacRudVLKEQvbE63KEBqbnCJc36wFiEyh3JkYw8e894i-mdZ4i86qn8c0TFP5qBcsBK9Hs0cPQO4JE9QwM_-PGZQUp-vjkUz8Wb55T1G7H9-6V6vM83hFZayI5t2nxzPZzgTzsN5Vvt4A4ZQCmAs4"

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) Handle(ctx context.Context, event event) (common.StepResponse, error) {
	var output common.StepResponse

	var quoteId = event.Data.QuoteID

	req, err := http.NewRequest(http.MethodGet, getEndpoint(quoteId), nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return output, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	req.Header.Set("policyholder-dob", "1989-03-23")
	req.Header.Set("Host", "localhost")

	client := &http.Client{}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	validationErr := validateTest(resp)

	if validationErr != nil {
		return output, validationErr
	}

	output.Payload = "OK"

	return output, nil

}

/**
* build the endpoint to get an mta quote
**/
func getEndpoint(quoteId string) string {

	//  "https://m3yktt0wkj.execute-api.eu-west-1.amazonaws.com/prod/quote/:quoteId"

	return fmt.Sprintf("%s/prod/quote/%s", domain, quoteId)

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
