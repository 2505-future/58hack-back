package usecase

import (
	"58-hack-api/pkg/server/infrustructure"
	"context"
)

type IMessageUsecase interface {
	Send(ctx context.Context, roomID string, message []byte) error
}

type MessageUsecase struct {
	di infrustructure.IDynamoDB
	wi infrustructure.IWebsocketInfrastructure
}

func NewMessageUsecase(di infrustructure.IDynamoDB, wi infrustructure.IWebsocketInfrastructure) *MessageUsecase {
	return &MessageUsecase{
		di: di,
		wi: wi,
	}
}

func (mu *MessageUsecase) Send(ctx context.Context, roomID string, message []byte) error {
	var connectionIDs []string
	err := mu.di.GetConnectionIDs(roomID, &connectionIDs)
	if err != nil {
		return err
	}

	for _, connectionID := range connectionIDs {
		err = mu.wi.Send(ctx, connectionID, message)
		if err != nil {
			return err
		}
	}

	return nil
}
