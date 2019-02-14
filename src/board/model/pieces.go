package board

import (
	. "../../core/helpers"
	. "../../core/model"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type BoardPiece struct {
	ID        string `sql:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Board     BoardModel `gorm:"foreignkey:BoardID"`
	BoardID   string     `sql:"index"`
	User      User `gorm:"foreignkey:UserID"`
	UserID    string     `sql:"index"`
}

func (board *BoardPiece) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	HandleError(err)

	scope.SetColumn("ID", id.String())
	return nil
}
