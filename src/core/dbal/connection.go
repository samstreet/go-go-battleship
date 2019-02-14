package core

import (
	. "../helpers"
	"github.com/jinzhu/gorm"
)

func InitialiseConnection() *gorm.DB {
	DB, err := gorm.Open("sqlite3", "test.db")
	HandleError(err)
	return DB
}
