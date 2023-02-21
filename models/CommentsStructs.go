package models

type Comments struct {
	CommentID      int `gorm:"primaryKey"`
	UserID         int
	PostID         int
	CommentContent string `gorm:"type:varchar(1000)"`
}
