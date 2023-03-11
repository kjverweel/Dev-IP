package repositories

import (
	"log"
	"main.go/models"
)

func NewPost(Post *models.Posts) error {
	err := db.Create(Post).Error
	if err != nil {
		return err
	} else {
		log.Println("NewPost.go:Post is aangemaakt :D")
	}
	return nil
}
