package repositories

import (
	"log"
	"main.go/models"
)

func NewUsers(newUser *models.Users) error {
	err := db.Create(newUser).Error
	if err != nil {
		return err
	} else {
		log.Println("User is aangemaakt :D")
	}
	return nil
}
