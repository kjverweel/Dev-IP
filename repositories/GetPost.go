package repositories

import (
	"log"
	"main.go/models"
)

func GetRecentPosts() ([]models.Posts, error) {
	var recentPosts []models.Posts

	for i := 0; i < 8; i++ {
		var post models.Posts
		result := db.Table("Posts").Select("post_content").Order("created_at desc").Limit(1).Offset(i).Find(&post)
		if result.Error != nil {
			return nil, result.Error
		}
		recentPosts = append(recentPosts, post)
	}
	log.Println(recentPosts)
	return recentPosts, nil
}
