package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChallengeIdResponse struct {
	ChallengeId string `json:"id"`
}

func Challenge() (cid string, err error) {
	client := &http.Client{}
	jsonBody := []byte(`{"to": "0999999999"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, _ := http.NewRequest("POST", "https://zero-testing.aviva.co.uk/login/api/challenge", bodyReader)
	req.Header.Set("authorization", "Basic YXZpdmE6aG94dG9uMTI=")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "text/plain")

	res, err := client.Do(req)

	if err != nil {
		return
	}
	defer res.Body.Close()

	result := ChallengeIdResponse{}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		err = fmt.Errorf("failed to decode response: %w", err)
		return
	}

	cid = result.ChallengeId

	return
}

type JwtResponse struct {
	Jwt string `json:"jwt"`
}

func Response(cid string) (jwt string, err error) {
	client := &http.Client{}

	// fmt.Sprintf("%s,%s", date, time)

	jsonBody := []byte(fmt.Sprintf(`{"code": "000000","id":"%s"}`, cid))
	bodyReader := bytes.NewReader(jsonBody)

	req, _ := http.NewRequest("POST", "https://zero-testing.aviva.co.uk/login/api/challenge/response", bodyReader)

	req.Header.Set("authorization", "Basic YXZpdmE6aG94dG9uMTI=")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	res, err := client.Do(req)

	if err != nil {
		return
	}
	defer res.Body.Close()

	result := JwtResponse{}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		err = fmt.Errorf("failed to decode response: %w", err)
		return
	}

	jwt = result.Jwt

	return
}
