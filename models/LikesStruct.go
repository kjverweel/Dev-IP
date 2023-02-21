package models

type Likes struct {
	LikeID int `gorm:"primaryKey"`
	UserID int
	PostID int
}
