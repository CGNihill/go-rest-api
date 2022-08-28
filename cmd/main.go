package main

import (
	"log"

	todo "github.com/CGNihill/go-rest-api"
	"github.com/CGNihill/go-rest-api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runing http server: %s", err.Error())
	}
}
