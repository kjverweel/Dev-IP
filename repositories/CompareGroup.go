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

func CompareGroupname() {

}