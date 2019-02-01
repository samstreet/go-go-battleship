package model

import (
	BoardModels "../../board/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
)

type SessionModel struct {
	gorm.Model
	UUID  uuid.UUID
	Board BoardModels.BoardModel
}
