package infrustructure

import (
	"58-hack-api/pkg/server/model"
	"encoding/json"

	"github.com/supabase-community/supabase-go"
)

type IRoomInfrustructure interface {
	CreateRoom(hostID, name string, capacity int) (string, error)
	GetRoomByID(id string, room *model.Room) error
}

type RoomInfrustructure struct {
	client *supabase.Client
}

func NewRoomInfrustructure(client *supabase.Client) IRoomInfrustructure {
	return &RoomInfrustructure{
		client: client,
	}
}

func (ri *RoomInfrustructure) CreateRoom(hostID, name string, capacity int) (string, error) {
	resp, _, err := ri.client.
		From("rooms").
		Insert(map[string]interface{}{
			"host_id":  hostID,
			"name":     name,
			"capacity": capacity,
		}, false, "", "", "").
		Execute()
	if err != nil {
		return "", err
	}

	var room []model.Room

	if err := json.Unmarshal(resp, &room); err != nil {
		return "", err
	}

	return room[0].ID, nil
}

func (ri *RoomInfrustructure) GetRoomByID(id string, room *model.Room) error {
	resp, _, err := ri.client.
		From("rooms").
		Select("*", "exact", false).
		Eq("id", id).
		Single().
		Execute()

	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, room)
	if err != nil {
		return err
	}

	return nil
}
