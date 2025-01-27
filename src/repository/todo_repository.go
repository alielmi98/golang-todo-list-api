package repository

import (
	"context"
	"log"

	"github.com/alielmi98/golang-todo-list-api/constants"
	"github.com/alielmi98/golang-todo-list-api/data/db"
	"github.com/alielmi98/golang-todo-list-api/data/models"

	"gorm.io/gorm"
)

type ToDoRepository interface {
	CreateToDo(ctx context.Context, model *models.Todo) error
	UpdateToDo(ctx context.Context, id int, model *models.Todo) error
	DeleteToDo(ctx context.Context, id int) error
	GetToDoById(ctx context.Context, id int) (*models.Todo, error)
}

type toDoRepository struct {
	database *gorm.DB
}

func NewToDoRepository() ToDoRepository {
	return &toDoRepository{
		database: db.GetDb(),
	}
}

func (r *toDoRepository) CreateToDo(ctx context.Context, model *models.Todo) error {
	tx := r.database.WithContext(ctx).Begin()
	err := tx.
		Create(model).
		Error
	if err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Insert, err.Error())
		return err
	}
	tx.Commit()

	return nil
}

func (r *toDoRepository) UpdateToDo(ctx context.Context, id int, model *models.Todo) error {
	tx := r.database.WithContext(ctx).Begin()
	if err := tx.Model(model).
		Where("id = ? ", id).
		Updates(model).
		Error; err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Update, err.Error())
		return err
	}
	tx.Commit()

	return nil
}

func (r *toDoRepository) DeleteToDo(ctx context.Context, id int) error {
	tx := r.database.WithContext(ctx).Begin()

	model := new(models.Todo)
	if err := tx.Where("id = ?", id).Delete(model).Error; err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Delete, err.Error())
		return err
	}
	tx.Commit()

	return nil
}
func (r *toDoRepository) GetToDoById(ctx context.Context, id int) (*models.Todo, error) {
	model := new(models.Todo)
	err := r.database.
		Where("id = ? ", id).
		First(model).
		Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
