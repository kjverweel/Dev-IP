package repositories

import (
	"github.com/jinzhu/gorm"
	"log"
	"main.go/models"
)

func CheckGroupMembers(Groupmembers *models.Groupmembers) (bool, error) {
	err := db.Where("user_id = ?", Groupmembers.UserID, "groep_id = ?", Groupmembers.GroepID).Error
	if err == gorm.ErrRecordNotFound {
		return false, err
	} else if err != nil {
		log.Println("this is an error")
	}

	return true, nil
}
