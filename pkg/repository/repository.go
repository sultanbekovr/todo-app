package repository

import "github.com/sultanbekovr/todo-app/pkg/repository"

type Authorization struct {
}

type TodoList struct {
}

type TodoItem struct {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func newRepository(repos *repository.Repository) *Repository {
	return &Service{}
}
