package main

import (
	"github.com/sultanbekovr/todo-app"
	"github.com/sultanbekovr/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}
}
