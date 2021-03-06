package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// MyEvent -
type MyEvent struct {
	IsTest bool   `json:"isTest"`
	Name   string `json:"name"`
}

// HandleRequest -
func HandleRequest(ctx context.Context, event MyEvent) (MyEvent, error) {
	fmt.Println("event : ", event)
	return event, nil
}

func main() {
	lambda.Start(HandleRequest)
}
