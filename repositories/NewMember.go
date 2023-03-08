package repositories

import (
	"log"
	"main.go/models"
)

func NewMember(newMember *models.Groupmembers, groups *models.Groups) error {

	err := db.Create(newMember).Error
	if err != nil {
		return err
	} else {
		log.Println("Lid is toegevoegd aan de groep")
	}
	return nil
}
