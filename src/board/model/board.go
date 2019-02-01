package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
)

type BoardModel struct {
	gorm.Model
	UUID uuid.UUID
}
