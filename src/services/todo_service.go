package services

import (
	"context"

	"github.com/alielmi98/golang-todo-list-api/api/dto"
	"github.com/alielmi98/golang-todo-list-api/constants"
	"github.com/alielmi98/golang-todo-list-api/data/models"
	"github.com/alielmi98/golang-todo-list-api/repository"
)

type ToDoService interface {
	CreateToDo(ctx context.Context, todo *dto.CreateTodoRequest) (*dto.TodoResponse, error)
}

type toDoService struct {
	todoRepository repository.ToDoRepository
}

func NewToDoService() ToDoService {
	return &toDoService{
		todoRepository: repository.NewToDoRepository(),
	}
}

func (s *toDoService) CreateToDo(ctx context.Context, todo *dto.CreateTodoRequest) (*dto.TodoResponse, error) {
	userId := int(ctx.Value(constants.UserIdKey).(float64))
	model := models.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		UserId:      userId,
	}
	response, err := s.todoRepository.CreateToDo(ctx, &model)

	return response, err
}
