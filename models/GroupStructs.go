package models

import "gorm.io/gorm"

type Groups struct {
	gorm.Model
	groep_name  string `gorm:"type:varchar(255)"`
	groep_admin string `gorm:"type:varchar(255)"`
}
