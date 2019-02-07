package model

import (
	"../dbal"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	SessionID  string
	Connection *gorm.DB
}

func (user User) NewUser() *User {
	return &User{Connection: dbal.InitialiseConnection()}
}

func (user *User) Fresh() *User {
	return user.NewUser()
}

func (user User) GetID() uuid.UUID {
	id, _ := uuid.FromString(user.ID)
	return id
}
