package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"main.go/models"
)

func CheckGroup(newGroup *models.Groups) (bool, error) {
	var existingGroup models.Groups
	err := db.Where("groepname = ?", newGroup.Groepname).First(&existingGroup).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		fmt.Println("This is where it go wrong wrong")
	}
	return true, nil
}
