package service

import "github.com/CGNihill/go-rest-api/pkg/repository"

type Authorization interface{}
type TodoList interface{}
type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}