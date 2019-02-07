package model

import (
	"../../board/model"
	"../../core/dbal"
	"../../core/helpers"
	CoreModels "../../core/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
	"time"
)

type SessionModel struct {
	ID         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Board      model.BoardModel `gorm:"foreignkey:BoardID"`
	BoardID    string
	Users      []CoreModels.User
	Connection *gorm.DB
}

func (session *SessionModel) NewSessionModel() *SessionModel {
	return &SessionModel{Connection: dbal.InitialiseConnection()}
}

func (session *SessionModel) Fresh() *SessionModel {
	return session.NewSessionModel()
}

func (session *SessionModel) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	helpers.HandleError(err)

	scope.SetColumn("ID", id.String())
	return nil
}
