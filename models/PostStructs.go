package models

import "gorm.io/gorm"

type Posts struct {
	gorm.Model
	GroepID     int
	PostContent string `gorm:"type:varchar(1000)"`
	UserID      int
}
