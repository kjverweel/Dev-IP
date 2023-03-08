package repositories

import (
	"gorm.io/gorm"
	"log"
	"main.go/models"
)

func CheckGroup(newGroup *models.Groups) (bool, error) {
	var existingGroup models.Groups
	err := db.Where("groepname = ?", newGroup.Groepname).First(&existingGroup).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		log.Println("CompareGroup.go:This is where it go wrong wrong")
	}
	return true, nil
}

func CompareGroupname(GetID *models.Groups) (int, error) {
	var GroupID int
	err := db.Model(&models.Groups{}).Select("id").Where("groepname = ?", GetID.Groepname).Scan(&GroupID).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CompareUsers.go:This is a database fault")
		return 0, err
	} else if err != nil {
		log.Println("CompareUsers.go:Probably couldn't find ID")
		return 0, err
	}
	return GroupID, nil
}
