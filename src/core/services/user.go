package services

import (
	"../model"
	"errors"
)

type UserService struct {
	Model model.User
}

func NewUserService() *UserService {
	m := model.User{}
	return &UserService{Model: *m.Fresh()}
}

func (service UserService) FindByUUID(uuid string) (user *model.User, err error)  {
	tmp := service.Model.Fresh()
	service.Model.Connection.Where("id = ?", uuid).First(&tmp)
	if tmp.GetID().String() == "00000000-0000-0000-0000-000000000000"{
		err = errors.New("invalid user")
	}

	user = tmp
	return user, err
}
