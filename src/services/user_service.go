package services

import (
	"log"

	"github.com/alielmi98/golang-todo-list-api/api/dto"
	"github.com/alielmi98/golang-todo-list-api/config"
	"github.com/alielmi98/golang-todo-list-api/constants"

	"github.com/alielmi98/golang-todo-list-api/data/db"
	"github.com/alielmi98/golang-todo-list-api/data/models"

	"github.com/alielmi98/golang-todo-list-api/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type UserService struct {
	cfg          *config.Config
	tokenService *TokenService
	database     *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	return &UserService{
		cfg:          cfg,
		database:     database,
		tokenService: NewTokenService(cfg),
	}
}

// Login by username
func (s *UserService) LoginByUsername(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var user models.User
	err := s.database.
		Model(&models.User{}).
		Where("username = ?", req.Username).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	tdto := tokenDto{UserId: user.Id, Email: user.Email}

	token, err := s.tokenService.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil

}

// Register by username
func (s *UserService) RegisterByUsername(req *dto.RegisterUserByUsernameRequest) error {
	u := models.User{Username: req.Username, Email: req.Email}

	exists, err := s.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = s.existsByUsername(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.General, constants.HashPassword, err.Error())
		return err
	}
	u.Password = string(hp)

	tx := s.database.Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Rollback, err.Error())
		return err
	}
	tx.Commit()
	return nil

}

func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Select, err.Error())
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Select, err.Error())
		return false, err
	}
	return exists, nil
}
