package usecase

import (
	"58-hack-api/pkg/server/infrustructure"
	"58-hack-api/pkg/server/model"
)

type IRoomUsecase interface {
	CreateRoom(hostID, name string, capacity int) (string, error)
	Verify(id string) error
	JoinRoom(id string) ([]model.User, error)
}

type RoomUsecase struct {
	ri infrustructure.IRoomInfrustructure
	di infrustructure.IDynamoDB
}

func NewRoomUsecase(ri infrustructure.IRoomInfrustructure, di infrustructure.IDynamoDB) IRoomUsecase {
	return &RoomUsecase{
		ri: ri,
		di: di,
	}
}

func (ru *RoomUsecase) CreateRoom(hostID, name string, capacity int) (string, error) {
	return ru.ri.CreateRoom(hostID, name, capacity)
}

func (ru *RoomUsecase) Verify(id string) error {
	room := &model.Room{}
	err := ru.ri.GetRoomByID(id, room)
	if err != nil {
		return err
	}
	return nil
}

func (ru *RoomUsecase) JoinRoom(id string) ([]model.User, error) {
	users := []model.User{}
	err := ru.di.GetUsers(id, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
