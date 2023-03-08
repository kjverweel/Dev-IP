package repositories

import (
	"github.com/jinzhu/gorm"
	"log"
	"main.go/models"
)

func GetGroup(groepname string) ([]models.Groups, error) {
	var groups []models.Groups
	err := db.Select("groepname", groepname).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // return nil if no records found
		}
		return nil, err
	}
	log.Println(groups)
	return groups, nil
}
