package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	UserNickname string `gorm:"type:varchar(255)"`
	UserEmail    string `gorm:"type:varchar(255)"`
	UserPassword string `gorm:"type:varchar(255)"`
}
