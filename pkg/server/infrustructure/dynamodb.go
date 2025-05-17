package infrustructure

import (
	"58-hack-api/pkg/server/model"
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type IDynamoDB interface {
	GetConnectionIDs(roomID string, connectionIDs *[]string) error
	GetUsers(roomId string, users *[]model.User) error
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

func (d *DynamoDB) GetUsers(roomId string, users *[]model.User) error {
	output, err := d.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(d.tableName),
		IndexName:              aws.String("roomId-index"),
		KeyConditionExpression: aws.String("roomId = :roomId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":roomId": &types.AttributeValueMemberS{Value: roomId},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to query by roomId: %w", err)
	}

	for _, item := range output.Items {
		user := model.User{}
		if val, ok := item["userId"].(*types.AttributeValueMemberS); ok {
			user.UserID = val.Value
		}
		if val, ok := item["iconUrl"].(*types.AttributeValueMemberS); ok {
			user.IconUrl = val.Value
		}
		if val, ok := item["cd"].(*types.AttributeValueMemberN); ok {
			user.Cd, err = strconv.Atoi(val.Value)
			if err != nil {
				return fmt.Errorf("failed to convert cd to int: %w", err)
			}
		}
		if val, ok := item["power"].(*types.AttributeValueMemberN); ok {
			user.Power, err = strconv.Atoi(val.Value)
			if err != nil {
				return fmt.Errorf("failed to convert power to int: %w", err)
			}
		}
		if val, ok := item["weight"].(*types.AttributeValueMemberN); ok {
			user.Weight, err = strconv.Atoi(val.Value)
			if err != nil {
				return fmt.Errorf("failed to convert weight to int: %w", err)
			}
		}
		if val, ok := item["volume"].(*types.AttributeValueMemberN); ok {
			user.Volume, err = strconv.Atoi(val.Value)
			if err != nil {
				return fmt.Errorf("failed to convert volume to int: %w", err)
			}
		}
		*users = append(*users, user)
	}

	return nil
}
