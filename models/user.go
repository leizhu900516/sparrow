package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name string
	Age uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}


type UserInfo struct {
	Username string 	`form:"username" json:"username"`
	Password string 	`form:"password" json:"password"`
}

//Db
