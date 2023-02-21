package repositories

import (
	"fmt"
	"main.go/models"
)

func NewUsers(newUser *models.Users) error {
	err := db.Create(newUser).Error
	if err != nil {
		return err
	} else {
		fmt.Println("User is aangemaakt :D")
	}
	return nil
}
