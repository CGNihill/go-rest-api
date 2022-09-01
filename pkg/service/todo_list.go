package service

import (
	todo "github.com/CGNihill/go-rest-api"
	"github.com/CGNihill/go-rest-api/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func newTodoListService(r repository.TodoList) *TodoListService {
	return &TodoListService{r}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
