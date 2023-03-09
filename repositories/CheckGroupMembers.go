package repositories

import "main.go/models"

func CheckGroupMembers(Groupmembers *models.Groupmembers) (bool, error) {
	err := db.Where("user_id = ?", Groupmembers.UserID, "groep_id = ?", Groupmembers.GroepID).Error
	return false, nil
}
