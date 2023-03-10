package repositories

import (
	"log"
	"main.go/models"
)

func CheckGroupMembers(Groupmembers *models.Groupmembers) (bool, error) {
	var member models.Groupmembers
	err := db.Where("user_id = ?", Groupmembers.UserID).Where("groep_id", Groupmembers.GroepID).First(&member).Error
	log.Println(member)
	if err != nil {
		log.Println("checkgroupmembers:", err)
		return true, nil
	}
	return false, nil
}
