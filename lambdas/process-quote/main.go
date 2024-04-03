package main

import (
	// "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-lambda-go/lambda"
)

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
