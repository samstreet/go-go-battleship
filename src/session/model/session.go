package model

import (
	"../../board/model"
	"../../core/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
	"time"
)

type SessionModel struct {
	ID string
	CreatedAt time.Time
	UpdatedAt time.Time
	Board     model.BoardModel
	BoardID string
}

func NewSessionModel() *SessionModel {
	return &SessionModel{}
}

func (session *SessionModel) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	helpers.HandleError(err)

	scope.SetColumn("ID", id.String())
	return nil
}
