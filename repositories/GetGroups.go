package repositories

import (
	"github.com/jinzhu/gorm"
	"log"
	"main.go/models"
)

func GetGroup() ([]models.Groups, error) {
	var groups []models.Groups
	err := db.Find(&groups).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // return nil if no records found
		}
		return nil, err
	}
	log.Println(groups)
	return groups, nil
}
