package board

import (
	. "../../core/helpers"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type ShipModel struct {
	ID     string `sql:"index"`
	Name   string
	Length int
}

func (board *ShipModel) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	HandleError(err)

	scope.SetColumn("ID", id.String())
	return nil
}
