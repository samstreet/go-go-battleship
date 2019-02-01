package dbal

import (
	"../helpers"
	"github.com/jinzhu/gorm"
)

func InitialiseConnection() *gorm.DB {
	DB, error := gorm.Open("sqlite3", "test.db")
	helpers.HandleError(error)
	return DB
}
