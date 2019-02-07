package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	Connection *gorm.DB
}

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
