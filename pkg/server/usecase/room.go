package usecase

import (
	"58-hack-api/pkg/server/infrustructure"
	"58-hack-api/pkg/server/model"
)

type IRoomUsecase interface {
	CreateRoom(hostID, name string, capacity int) (string, error)
	Verify(id string) error
}

type RoomUsecase struct {
	ri infrustructure.IRoomInfrustructure
}

func NewRoomUsecase(ri infrustructure.IRoomInfrustructure) IRoomUsecase {
	return &RoomUsecase{
		ri: ri,
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
