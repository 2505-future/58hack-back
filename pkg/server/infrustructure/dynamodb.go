package infrustructure

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type IDynamoDB interface {
	GetConnectionIDs(roomID string, connectionIDs *[]string) error
}

type DynamoDB struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoDB(client *dynamodb.Client, tableName string) *DynamoDB {
	return &DynamoDB{
		client:    client,
		tableName: tableName,
	}
}

func (d *DynamoDB) GetConnectionIDs(roomID string, connectionIDs *[]string) error {
	output, err := d.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(d.tableName),
		IndexName:              aws.String("roomId-index"),
		KeyConditionExpression: aws.String("roomId = :roomId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":roomId": &types.AttributeValueMemberS{Value: roomID},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to query by roomId: %w", err)
	}

	for _, item := range output.Items {
		if val, ok := item["connectionId"].(*types.AttributeValueMemberS); ok {
			*connectionIDs = append(*connectionIDs, val.Value)
		}
	}

	return nil
}
