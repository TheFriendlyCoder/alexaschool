package main

import (
	"context"

	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

type Connection struct{}

func (connection Connection) IntentDispatcher(ctx context.Context, request alexa.Request) (alexa.Response, error) {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "HelloWorldIntent":
		response = alexa.NewSimpleResponse("Hi there from Kevin", "It is a wonderful day in the neighborhood")
	default:
		response = alexa.NewSimpleResponse("Unknown Request", "The intent was unrecognized")
	}
	return response, nil
}

func main() {
	connection := Connection{}
	lambda.Start(connection.IntentDispatcher)
}
