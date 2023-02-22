package repositories

import (
	"fmt"
	"main.go/models"
)

func NewGroup(newGroup *models.Groups) error {
	err := db.Create(newGroup).Error
	if err != nil {
		return err
	} else {
		fmt.Println("groep is aangemaakt :D")
	}
	return nil
}
