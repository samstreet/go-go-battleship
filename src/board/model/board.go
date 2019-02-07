package model

import (
	"../../core/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
	"time"
)

type BoardModel struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	XLength   int
	YLength   int
	Pieces    []BoardPiece
}

func (board *BoardModel) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	helpers.HandleError(err)

	scope.SetColumn("ID", id.String())
	return nil
}

type BoardPiece struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Board BoardModel `gorm:"foreignkey:BoardID"`
	BoardID string
}

func (board *BoardPiece) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	helpers.HandleError(err)

	scope.SetColumn("ID", id.String())
	return nil
}