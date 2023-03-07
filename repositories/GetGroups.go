package repositories

import (
	"github.com/jinzhu/gorm"
	"main.go/models"
)

func GetGroup(groupname string) ([]models.Groups, error) {
	var groups []models.Groups
	err := db.Where("groupname = ?", groupname).Find(&groups).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // return nil if no records found
		}
		return nil, err
	}
	return groups, nil
}
