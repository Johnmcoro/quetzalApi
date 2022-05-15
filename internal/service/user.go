package service

import (
	"log"

	"github.com/johnmcoro/quetzalapi/internal/models"
	"github.com/johnmcoro/quetzalapi/internal/storage/postgres"
)

type UserService interface {
	GetUsers() ([]models.UserDBModel, error)
}
type userService struct {
	Storage postgres.UserStorage
}

func NewUserService(userStorage postgres.UserStorage) *userService {
	return &userService{
		Storage: userStorage,
	}
}

func (us userService) GetUsers() ([]models.UserDBModel, error) {
	users, err := us.Storage.GetUsers()
	log.Println("service get users")
	if err != nil {
		log.Println("service get users error", err)

		return nil, err
	}
	return users, nil
}
