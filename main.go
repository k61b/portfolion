package main

import (
	"log"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/storage"
	"github.com/kayraberktuncer/portfolion/pkg/handlers"
)

func main() {
	store, err := storage.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	handlers := handlers.NewHandlers(":"+lib.GoDotEnvVariable("PORT"), store)

	handlers.Run()
}
