package repository

import (
	"context"
	"log"

	"github.com/alielmi98/golang-todo-list-api/api/dto"
	"github.com/alielmi98/golang-todo-list-api/common"
	"github.com/alielmi98/golang-todo-list-api/constants"
	"github.com/alielmi98/golang-todo-list-api/data/db"
	"github.com/alielmi98/golang-todo-list-api/data/models"

	"gorm.io/gorm"
)

type ToDoRepository interface {
	CreateToDo(ctx context.Context, model *models.Todo) (*dto.TodoResponse, error)
}

type toDoRepository struct {
	database *gorm.DB
}

func NewToDoRepository() ToDoRepository {
	return &toDoRepository{
		database: db.GetDb(),
	}
}

func (r *toDoRepository) CreateToDo(ctx context.Context, model *models.Todo) (*dto.TodoResponse, error) {
	tx := r.database.WithContext(ctx).Begin()
	err := tx.
		Create(model).
		Error
	if err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Insert, err.Error())
		return &dto.TodoResponse{}, err
	}
	tx.Commit()
	response, err := common.TypeConverter[dto.TodoResponse](model)
	if err != nil {
		return &dto.TodoResponse{}, err
	}
	return response, nil
}
