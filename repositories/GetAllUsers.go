package repositories

import (
	"github.com/jinzhu/gorm"
	"log"
)

func GetAllUsers() ([]string, error) {
	var AllUsers []string
	err := db.Table("users").Pluck("DISTINCT user_nickname", &AllUsers).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // return nil if no records found
		}
		return nil, err
	}
	log.Println(AllUsers)
	return AllUsers, nil
}
