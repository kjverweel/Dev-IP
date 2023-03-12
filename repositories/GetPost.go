package repositories

import (
	"main.go/models"
)

func GetPost() ([]models.Posts, error) {
	var posts []models.Posts
	result := db.Table("Posts").Select("post_content").Order("created_at desc").Limit(8).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
