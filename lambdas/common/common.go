package common

type StepOutput[T any] struct {
	Payload T       `json:"data"`
	Error   *string `json:"error"`
}

type StepInput[T any] struct {
	UserInfo struct {
		JWT string `json:"jwt"`
		Dob string `json:"dob"`
	} `json:"userInfo"`
	Host struct {
		Domain string `json:"domain"`
	} `json:"host"`
	Payload T `json:"data"`
}

type QuoteResponseBody struct {
	QuoteID       string `json:"quoteId"`
	QuoteRejected bool   `json:"quoteRejected"`
}
