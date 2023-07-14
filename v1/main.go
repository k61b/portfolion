package main

import (
	"log"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/storage"
	"github.com/kayraberktuncer/portfolion/pkg/handlers"
)

// @title Portfolion API
// @description This is a sample server Portfolion server.
// @version 0.1
// @host localhost:6161
// @BasePath /api/v1
func main() {
	store, err := storage.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	handlers := handlers.NewHandlers(":"+lib.GoDotEnvVariable("PORT"), store)

	lib.Logger()

	handlers.Run()
}
