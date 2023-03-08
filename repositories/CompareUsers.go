package repositories

import (
	"gorm.io/gorm"
	"log"
	"main.go/models"
)

func CompareUsers(newUser *models.Users) (bool, error) {
	var existingUser models.Users
	err := db.Where("user_nickname = ?", newUser.UserNickname).First(&existingUser).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		log.Println("CompareUsers.go:This is where it go wrong wrong")
	}
	return true, nil
}

func CompareUsername(GetID *models.Users) (int, error) {
	var UserID int
	err := db.Model(&models.Users{}).Select("id").Where("user_nickname = ?", GetID.UserNickname).Scan(&UserID).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CompareUsers.go:This is a database fault")
		return 0, err
	} else if err != nil {
		log.Println("CompareUsers.go:Probably couldn't find ID")
		return 0, err
	}
	return UserID, nil
}
