package repositories

import (
	"log"
	"main.go/models"
)

func CheckIfInGroup(UserID int, GroupID int) (bool, error) {
	var member models.Groupmembers
	err := db.Where("user_id = ? AND groep_id = ?", UserID, GroupID).First(&member).Error
	if err != nil {
		log.Println("checkgroupmembers:", err)
		return false, nil
	}
	return true, nil
}
