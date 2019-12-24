package main

import (
	"log"

	"github.com/hvs-fasya/facade/internal/models"
	nfier "github.com/hvs-fasya/facade/internal/notifier"
	"github.com/hvs-fasya/facade/internal/service"
)

func main() {
	n := nfier.NewNotifier()
	s := service.NewUserService(n)
	var newUsers = [2]models.CreateUserInput{
		{Name: "new_user"},
		{Name: ""},
	}
	for _, u := range newUsers {
		if err := s.CreateUser(u); err != nil {
			log.Printf("create user '" + u.Name + "' error: " + err.Error())
		}
	}
}
