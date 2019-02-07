package model

import (
	"../../board/model"
	"../../core/dbal"
	CoreModels "../../core/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

type SessionModel struct {
	ID         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Board      model.BoardModel `gorm:"foreignkey:BoardID"`
	BoardID    string
	Players    []CoreModels.User
	Connection *gorm.DB
}

func (session *SessionModel) NewSessionModel() *SessionModel {
	return &SessionModel{Connection: dbal.InitialiseConnection()}
}

func (session *SessionModel) Fresh() *SessionModel {
	return session.NewSessionModel()
}

func (session *SessionModel) BeforeCreate(scope *gorm.Scope) error {
	id, _ := uuid.NewV4()
	log.Fatal(scope.SetColumn("ID", id.String()))
	return nil
}
