package core

import (
	. "../dbal"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	SessionID  string
	PlayerRole int
	Connection *gorm.DB
}

func (user User) NewUser() *User {
	return &User{Connection: InitialiseConnection()}
}

func (user *User) Fresh() *User {
	return user.NewUser()
}

func (user User) GetID() uuid.UUID {
	id, _ := uuid.FromString(user.ID)
	return id
}