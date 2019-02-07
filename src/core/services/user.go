package services

import (
	"../model"
	u "github.com/satori/go.uuid"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (service UserService) findByUUID(uuid u.UUID) model.User {
	return model.User{ID: uuid.String()}
}
