package router

import (
	"58-hack-api/pkg/middleware"
	"58-hack-api/pkg/server/controller"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func NewRouter(rc controller.IRoomController, mc controller.IMessageController) *echo.Echo {
	e := echo.New()

	// panicが発生した場合の処理
	e.Use(echomiddleware.Recover())

	// CORSの設定
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		Skipper:      echomiddleware.DefaultCORSConfig.Skipper,
		AllowOrigins: echomiddleware.DefaultCORSConfig.AllowOrigins,
		AllowMethods: echomiddleware.DefaultCORSConfig.AllowMethods,
		AllowHeaders: []string{"Content-Type,Accept,Origin,x-token"},
	}))

	zapLogger, _ := zap.NewProduction()

	// 独自のログミドルウェアを追加
	e.Use(middleware.LoggingMiddleware(*zapLogger))

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	roomAPI := e.Group("/rooms")
	roomAPI.POST("", rc.CreateRoom)
	roomAPI.POST("/verify", rc.VerifyPassword)
	roomAPI.GET("/:roomID/join", rc.JoinRoom)
	roomAPI.POST("/:roomID/start", mc.SendStart)
	roomAPI.POST("/:roomID/action", mc.SendAction)

	return e
}
