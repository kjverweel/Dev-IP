package repositories

import (
	"main.go/models"
)

func GetRecentPosts() ([][]string, error) {
	var recentPostContents [][]string

	for i := 0; i < 8; i++ {
		var post models.Posts
		result := db.Table("Posts").Select("post_content").Order("created_at desc").Limit(1).Offset(i).Find(&post)

		if result.Error != nil {
			return nil, result.Error
		}
		recentPostContents = append(recentPostContents, []string{post.PostContent})
	}

	return recentPostContents, nil
}
