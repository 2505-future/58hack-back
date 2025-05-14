package main

import (
	"58-hack-api/pkg/server/controller"
	"58-hack-api/pkg/server/db"
	"58-hack-api/pkg/server/infrustructure"
	"58-hack-api/pkg/server/router"
	"58-hack-api/pkg/server/usecase"
)

func main() {
	db := db.NewClient()
	ri := infrustructure.NewRoomInfrustructure(db)
	ru := usecase.NewRoomUsecase(ri)
	rc := controller.NewRoomController(ru)

	e := router.NewRouter(rc)

	e.Logger.Fatal(e.Start(":8080"))
}
