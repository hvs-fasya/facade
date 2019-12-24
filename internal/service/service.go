package service

import (
	"errors"
	"sync"

	"github.com/hvs-fasya/facade/internal/models"
)

type UserServiceIface interface {
	CreateUser(input models.CreateUserInput) error
}

type NotifierIface interface {
	NotifyCreateUser(string)
}

type userService struct {
	mux      sync.Mutex
	storage  storage
	notifier NotifierIface
}

func (s *userService) CreateUser(inp models.CreateUserInput) error {
	var curID int
	if inp.Name == "" {
		return errors.New("user name required")
	}
	s.mux.Lock()
	for _, u := range s.storage {
		if curID < u.ID {
			curID = u.ID
		}
	}
	s.storage = append(s.storage, models.User{ID: curID + 1, Name: inp.Name})
	s.mux.Unlock()
	//hidden behind facade action
	s.notifier.NotifyCreateUser(inp.Name)

	return nil
}

type storage []models.User

func NewUserService(n NotifierIface) UserServiceIface {
	s := make([]models.User, 0)
	return &userService{
		storage:  s,
		notifier: n,
	}
}
