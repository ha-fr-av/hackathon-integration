package main

import (
	"context"
	"fmt"
)

func main() {

	err := handler(context.Background(), map[string]any{})
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}
