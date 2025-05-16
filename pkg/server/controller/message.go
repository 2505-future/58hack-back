package controller

import (
	"58-hack-api/pkg/server/usecase"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IMessageController interface {
	SendAction(c echo.Context) error
	SendJoin(c echo.Context) error
	SendLeave(c echo.Context) error
	SendStart(c echo.Context) error
}

type MessageController struct {
	uc usecase.IMessageUsecase
}

func NewMessageController(uc usecase.IMessageUsecase) *MessageController {
	return &MessageController{
		uc: uc,
	}
}

func (mc *MessageController) SendStart(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to read request body"})
	}
	return mc.sendRawJSON(c, body)
}

func (mc *MessageController) SendAction(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to read request body"})
	}
	return mc.sendRawJSON(c, body)
}

func (mc *MessageController) sendRawJSON(c echo.Context, raw []byte) error {
	roomID := c.Param("roomID")

	if err := mc.uc.Send(c.Request().Context(), roomID, raw); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusOK)
}
