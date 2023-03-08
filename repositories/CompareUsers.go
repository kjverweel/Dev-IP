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
	log.Println("getID:", GetID)
	var UserID int
	log.Println("UserID:", UserID)
	err := db.Table("users").Where("user_nickname = ?", GetID.UserNickname).First("id", &UserID).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CompareUsers.go:This is a database fault")
	} else if err != nil {
		log.Println("CompareUsers.go:Probably couldn't find ID")
	}
	log.Println(UserID)
	return UserID, nil
}
