package services

import (
	. "../model"
	"errors"
)

type UserService struct {
	Model User
}

func NewUserService() *UserService {
	m := User{}
	return &UserService{Model: *m.Fresh()}
}

func (service UserService) FindByUUID(uuid string) (user *User, err error)  {
	user = service.Model.Fresh()
	service.Model.Connection.Where("id = ?", uuid).First(&user)
	if user.ID == ""{
		err = errors.New("user not found")
	}

	return user, err
}
