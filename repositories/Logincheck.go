package repositories

import (
	"log"
	"main.go/models"
)

func LoginUser(YouExist *models.Users) bool {
	err := db.Where("user_nickname = ? AND user_password = ?", YouExist.UserNickname, YouExist.UserPassword).First(&YouExist).Error
	if err != nil {
		log.Println(err.Error())
		return false
	} else {
		log.Println("Logincheck.go:Login Succesvol")
	}

	return true
}
