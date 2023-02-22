package repositories

import (
	"log"
	"main.go/models"
)

func NewGroup(newGroup *models.Groups) error {
	err := db.Create(newGroup).Error
	if err != nil {
		return err
	} else {
		log.Println("groep is aangemaakt :D")
	}
	return nil
}
