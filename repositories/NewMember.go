package repositories

import (
	"log"
	"main.go/models"
)

func NewMember(Groupmembers *models.Groupmembers) error {
	err := db.Create(Groupmembers).Error
	if err != nil {
		return err
	} else {
		log.Println("NewMembers.go:", Groupmembers, "toegevoegd!")
	}
	return nil
}
