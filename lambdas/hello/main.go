package main

import (
	"context"
	"fmt"

	// "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-lambda-go/lambda"
)

type output struct {
	Hello string `json:"hello"`
}

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

type StepFunctionInput struct {
	Message string `json:"message"`
}

func (h Handler) Handle(ctx context.Context, event StepFunctionInput) (output, error) {

	fmt.Println("hello world")
	// fmt.Println(event)

	o := output{Hello: event.Message}

	return o, nil
}

func main() {

	h := NewHandler()
	lambda.Start(h.Handle)
}
