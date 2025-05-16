package infrustructure

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

type IWebsocketInfrastructure interface {
	Send(ctx context.Context, connectionID string, message []byte) error
}

type WebsocketInfrastructure struct{}

func NewWebsocketInfrastructure() *WebsocketInfrastructure {
	return &WebsocketInfrastructure{}
}

var endpoint = os.Getenv("APIGATEWAY_WEBSOCKET_ENDPOINT")

func (wi *WebsocketInfrastructure) Send(ctx context.Context, connectionID string, message []byte) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}

	client := apigatewaymanagementapi.NewFromConfig(cfg, func(o *apigatewaymanagementapi.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connectionID),
		Data:         message,
	}
	_, err = client.PostToConnection(ctx, input)

	log.Println("send message to connectionID:", connectionID)
	log.Println("message:", string(message))

	return err
}
