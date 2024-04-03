package main

import (
	"context"

	// "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-lambda-go/lambda"
)

type output struct{ Jwt string }

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Handle(ctx context.Context) (o output, err error) {

	cid, err := Challenge()

	if err != nil {
		return
	}
	jwt, err := Response(cid)

	if err != nil {
		return
	}

	return output{Jwt: jwt}, nil
}

func main() {

	h := NewHandler()
	lambda.Start(h.Handle)

	// fmt.Printf(data)
}
