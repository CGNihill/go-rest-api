package main

import (
	"log"

	todo "github.com/CGNihill/go-rest-api"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while runing http server: %s", err.Error())
	}
}
