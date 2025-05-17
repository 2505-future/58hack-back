package main

import (
	"58-hack-api/pkg/server/controller"
	"58-hack-api/pkg/server/db"
	"58-hack-api/pkg/server/infrustructure"
	"58-hack-api/pkg/server/router"
	"58-hack-api/pkg/server/usecase"
)

func main() {
	DBClient := db.NewClient()
	dynamodb := db.NewDynamoDBClient()
	ri := infrustructure.NewRoomInfrustructure(DBClient)
	di := infrustructure.NewDynamoDB(dynamodb, "websocket")
	wi := infrustructure.NewWebsocketInfrastructure()
	mu := usecase.NewMessageUsecase(di, wi)
	mc := controller.NewMessageController(mu)
	ru := usecase.NewRoomUsecase(ri, di)
	rc := controller.NewRoomController(ru)

	e := router.NewRouter(rc, mc)

	e.Logger.Fatal(e.Start(":8080"))
}
