package services

import (
	"../model"
)

type UserService struct {
	Model model.User
}

func NewUserService() *UserService {
	m := model.User{}
	return &UserService{Model: *m.Fresh()}
}

func (service UserService) FindByUUID(uuid string) *model.User {
	tmp := service.Model.Fresh()
	service.Model.Connection.Where("id = ?", uuid).First(&tmp)
	return tmp
}
