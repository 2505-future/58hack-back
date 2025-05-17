package schema

import "58-hack-api/pkg/server/model"

type CreateRoomRequest struct {
	HostID   string `json:"host_id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type CreateRoomResponse struct {
	ID     string `json:"room_id"`
	HostID string `json:"host_id"`
}

type VerifyRoomRequest struct {
	Password string `json:"password"`
}

type VerifyRoomResponse struct {
	Status string `json:"status"`
}

type JoinRoomResponse struct {
	Users []model.User `json:"users"`
}
