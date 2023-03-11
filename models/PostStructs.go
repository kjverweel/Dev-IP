package models

type Posts struct {
	PostID      int `gorm:"primaryKey"`
	GroepID     int
	PostContent string `gorm:"type:varchar(1000)"`
	UserID      int
}
