package services

import (
	"context"

	"github.com/alielmi98/golang-todo-list-api/api/dto"
	"github.com/alielmi98/golang-todo-list-api/common"
	"github.com/alielmi98/golang-todo-list-api/constants"
	"github.com/alielmi98/golang-todo-list-api/data/models"
	"github.com/alielmi98/golang-todo-list-api/repository"
	"github.com/mitchellh/mapstructure"
)

type ToDoService interface {
	CreateToDo(ctx context.Context, todo *dto.CreateToDoRequest) (*dto.ToDoResponse, error)
	UpdateToDo(ctx context.Context, id int, todo *dto.UpdateToDoRequest) (*dto.ToDoResponse, error)
	DeleteToDo(ctx context.Context, id int) error
	GetToDoById(ctx context.Context, id int) (*dto.ToDoResponse, error)
}

type toDoService struct {
	todoRepository repository.ToDoRepository
}

func NewToDoService() ToDoService {
	return &toDoService{
		todoRepository: repository.NewToDoRepository(),
	}
}

func (s *toDoService) CreateToDo(ctx context.Context, todo *dto.CreateToDoRequest) (*dto.ToDoResponse, error) {
	userId := int(ctx.Value(constants.UserIdKey).(float64))
	model := new(models.Todo)
	if err := mapstructure.Decode(todo, model); err != nil {
		return nil, err
	}
	model.UserId = userId

	err := s.todoRepository.CreateToDo(ctx, model)
	if err != nil {
		return &dto.ToDoResponse{}, err
	}

	response, err := common.TypeConverter[dto.ToDoResponse](model)
	if err != nil {
		return &dto.ToDoResponse{}, err
	}

	return response, err
}

func (s *toDoService) UpdateToDo(ctx context.Context, id int, todo *dto.UpdateToDoRequest) (*dto.ToDoResponse, error) {
	model := new(models.Todo)
	if err := mapstructure.Decode(todo, model); err != nil {
		return nil, err
	}

	err := s.todoRepository.UpdateToDo(ctx, id, model)
	if err != nil {
		return &dto.ToDoResponse{}, err
	}

	response, err := common.TypeConverter[dto.ToDoResponse](model)
	if err != nil {
		return &dto.ToDoResponse{}, err
	}
	return response, err

}
func (s *toDoService) GetToDoById(ctx context.Context, id int) (*dto.ToDoResponse, error) {

	model, err := s.todoRepository.GetToDoById(ctx, id)
	if err != nil {
		return &dto.ToDoResponse{}, err
	}

	response, err := common.TypeConverter[dto.ToDoResponse](model)
	if err != nil {
		return &dto.ToDoResponse{}, err
	}
	return response, err

}

func (s *toDoService) DeleteToDo(ctx context.Context, id int) error {
	err := s.todoRepository.DeleteToDo(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
