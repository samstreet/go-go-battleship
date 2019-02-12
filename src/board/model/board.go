package model

import (
	"../../core/dbal"
	"../../core/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
	"time"
)

type BoardModel struct {
	ID         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	XLength    int
	YLength    int
	Pieces     []BoardPiece
	Connection *gorm.DB
}

func NewBoardModel(connection *gorm.DB) *BoardModel {
	return &BoardModel{Connection: connection}
}

func (board *BoardModel) Fresh() *BoardModel {
	return NewBoardModel(dbal.InitialiseConnection())
}

func (board *BoardModel) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	helpers.HandleError(err)

	scope.SetColumn("ID", id.String())
	return nil
}
