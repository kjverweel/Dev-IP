package models

import "gorm.io/gorm"

type Groups struct {
	gorm.Model
	Groepname    string `gorm:"type:varchar(255);not null"`
	GroepadminID string `gorm:"type:varchar(100)"`
}
