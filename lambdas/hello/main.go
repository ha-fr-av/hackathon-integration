package main

import (
	"context"
	"fmt"

	// "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-lambda-go/lambda"
)

type output struct{ hello string }

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Handle(ctx context.Context) (output, error) {

	fmt.Println("hello world")

	o := output{hello: "world"}

	return o, nil
}

func main() {

	// _, err := Handler(context.Background())
	h := NewHandler()

	// if err != nil {
	// 	fmt.Printf("%s", err.Error())
	// }

	// lambda.start(h.Handle)
	lambda.Start(h.Handle)

	// fmt.Printf(data)
}
