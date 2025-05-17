package controller

import (
	"log"
	"net/http"

	"58-hack-api/pkg/server/controller/schema"
	"58-hack-api/pkg/server/usecase"

	"github.com/labstack/echo/v4"
)

type IRoomController interface {
	CreateRoom(c echo.Context) error
	VerifyPassword(c echo.Context) error
	JoinRoom(c echo.Context) error
}

type RoomController struct {
	ru usecase.IRoomUsecase
}

func NewRoomController(ru usecase.IRoomUsecase) IRoomController {
	return &RoomController{
		ru: ru,
	}
}

func (rc *RoomController) CreateRoom(c echo.Context) error {
	var req schema.CreateRoomRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	id, err := rc.ru.CreateRoom(req.HostID, req.Name, req.Capacity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	resp := schema.CreateRoomResponse{
		ID:     id,
		HostID: req.HostID,
	}

	return c.JSON(http.StatusCreated, resp)
}

func (rc *RoomController) VerifyPassword(c echo.Context) error {
	var req schema.VerifyRoomRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	err := rc.ru.Verify(req.Password)

	var resp schema.VerifyRoomResponse
	if err != nil {
		resp = schema.VerifyRoomResponse{
			Status: "ng",
		}
		return c.JSON(http.StatusNotFound, resp)
	}

	resp = schema.VerifyRoomResponse{
		Status: "ok",
	}

	return c.JSON(http.StatusOK, resp)
}

func (rc *RoomController) JoinRoom(c echo.Context) error {
	roomID := c.Param("roomID")
	log.Println("roomID:", roomID)
	if roomID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "roomID is required"})
	}

	users, err := rc.ru.JoinRoom(roomID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	resp := schema.JoinRoomResponse{
		Users: users,
	}

	return c.JSON(http.StatusOK, resp)
}
