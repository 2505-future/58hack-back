package main

import (
	"58-hack-api/pkg/server/router"
)

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))
}
