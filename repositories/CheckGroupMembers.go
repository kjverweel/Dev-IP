package repositories

import (
	"log"
	"main.go/models"
)

func CheckGroupMembers(Groupmembers *models.Groupmembers) (bool, error) {
	var member models.Groupmembers
	err := db.Where("user_id = ?", Groupmembers.UserID, "group_id", Groupmembers.GroepID).First(&member).Error
	log.Println(member)
	if err != nil {
		return false, nil
		if err.Error() == "record not found" {
			return false, nil
		}
		log.Println("Error querying database:", err)
		return false, err
	}
	return true, nil
}
