package repository

import (
	"context"
	"log"

	"github.com/alielmi98/golang-todo-list-api/api/dto"
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
	GetToDosByFilter(ctx context.Context, pagination *dto.PaginationInputWithFilter, userid int) (*dto.PagedList[dto.ToDoResponse], error)
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

func (r *toDoRepository) GetToDosByFilter(ctx context.Context, pagination *dto.PaginationInputWithFilter, userid int) (*dto.PagedList[dto.ToDoResponse], error) {
	var todos []models.Todo
	var totalRows int64

	offset := (pagination.PageNumber - 1) * pagination.PageSize

	query := r.database.Model(&models.Todo{})
	query = query.Where("user_id = ?", userid)

	for field, value := range pagination.Filter {
		switch v := value.(type) {
		case string:
			query = query.Where(field+" LIKE ?", "%"+v+"%")
		case bool:
			query = query.Where(field+" = ?", v)
		}

	}

	for field, order := range pagination.Sort {
		if order == "asc" {
			query = query.Order(field + " ASC")
		} else if order == "desc" {
			query = query.Order(field + " DESC")
		}
	}

	if err := query.Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if err := query.Offset(offset).Limit(pagination.PageSize).Find(&todos).Error; err != nil {
		return nil, err
	}

	var todoResponses []dto.ToDoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, dto.ToDoResponse{
			Id:          todo.Id,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			UserId:      todo.UserId,
		})
	}

	totalPages := int((totalRows + int64(pagination.PageSize) - 1) / int64(pagination.PageSize))
	hasNextPage := pagination.PageNumber < totalPages
	hasPrevPage := pagination.PageNumber > 1

	return &dto.PagedList[dto.ToDoResponse]{
		PageNumber:  pagination.PageNumber,
		PageSize:    pagination.PageSize,
		TotalRows:   totalRows,
		TotalPages:  totalPages,
		HasNextPage: hasNextPage,
		HasPrevPage: hasPrevPage,
		Items:       todoResponses,
	}, nil
}
