package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (string, error) {
	time.Sleep(time.Second * 1)
	return "Hello ƛ !", nil
}

func main() {
	lambda.Start(hello)
}
