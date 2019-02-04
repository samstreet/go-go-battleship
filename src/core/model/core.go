package model

import (
	SessionModels "../../session/model"
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	Connection *gorm.DB
}

type User struct {
	ID         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Session    SessionModels.SessionModel `gorm:"foreignkey:SessionID"`
	SessionID  string
	Model
}
