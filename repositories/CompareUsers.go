package repositories

import (
	"gorm.io/gorm"
	"main.go/models"
)

func CompareUsers(newUser *models.Users) (bool, error) {
	var existingUser models.Users
	err := db.Where("user_nickname = ?", newUser.UserNickname).First(&existingUser).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		log.Println("This is where it go wrong wrong")
	}
	return true, nil
}
