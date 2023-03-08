package repositories

import (
	"github.com/jinzhu/gorm"
	"log"
)

func GetGroup() ([]string, error) {
	var groups []string
	err := db.Table("groups").Pluck("DISTINCT groepname", &groups).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // return nil if no records found
		}
		return nil, err
	}
	log.Println(groups)
	return groups, nil
}
