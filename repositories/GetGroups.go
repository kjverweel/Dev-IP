package repositories

import (
	"github.com/jinzhu/gorm"
	"log"
	"main.go/models"
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
	if groups == nil {
		log.Println("GetGroups.go:groups is empty")
	} else {
		log.Println("GetGroups.go:Groups succesfully called")
	}
	return groups, nil
}

func GetGroups(GroupID []int) []string {
	var groups []string
	err := db.Table("groups").Where("id IN (?)", GroupID).Pluck("DISTINCT groepname", &groups).Error
	if err != nil {
		return nil
	}
	if groups == nil {
		log.Println("GetGroups.go:groups is empty")
	} else {
		log.Println("GetGroups.go:Groups succesfully called")
	}
	return groups
}

func GetLatestGroup() int {
	var latestGroup int
	err := db.Model(&models.Groups{}).Select("id").Table("groups").Order("created_at desc").First(&latestGroup).Error
	if err != nil {
		return 0
	}
	if latestGroup == 0 {
		log.Println("GetLatestGroup.go: no group found")
	} else {
		log.Println("GetLatestGroup.go: Group found")
	}
	return latestGroup
}
