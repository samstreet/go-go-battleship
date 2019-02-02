package model

import "github.com/jinzhu/gorm"

type Model struct {
	Connection *gorm.DB
}
