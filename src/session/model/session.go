package session

import (
	. "../../board/model"
	. "../../core/dbal"
	. "../../core/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
	"time"
)

type SessionModel struct {
	ID         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Board      BoardModel `gorm:"foreignkey:BoardID"`
	BoardID    string
	Players    []User `gorm:"ForeignKey:SessionID"`
	Connection *gorm.DB
}

func (session *SessionModel) NewSessionModel() *SessionModel {
	return &SessionModel{Connection: InitialiseConnection()}
}

func (session *SessionModel) Fresh() *SessionModel {
	return session.NewSessionModel()
}

func (session *SessionModel) BeforeCreate(scope *gorm.Scope) error {
	id, _ := uuid.NewV4()
	scope.SetColumn("ID", id.String())
	return nil
}
