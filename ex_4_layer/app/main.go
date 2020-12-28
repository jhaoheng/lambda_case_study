package main

import (
	"context"
	"plugin"

	"github.com/aws/aws-lambda-go/lambda"
)

func pluginActive() error {
	pluginModule, err := plugin.Open("/opt/layer.so")
	if err != nil {
		return err
	}
	HelloLayer1, err := pluginModule.Lookup("HelloLayer1")
	if err != nil {
		return err
	}
	HelloLayer1.(func())()
	return nil
}

// Handler -
func Handler(ctx context.Context) error {
	return pluginActive()
}

func main() {
	lambda.Start(Handler)
	// fmt.Println(pluginActive())
}
