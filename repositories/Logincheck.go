package repositories

import (
	"fmt"
	"main.go/models"
)

func LoginUser(YouExist *models.Users) bool {
	err := db.Where("user_nickname = ? AND user_password = ?", YouExist.UserNickname, YouExist.UserPassword).First(&YouExist).Error
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		fmt.Println("Login Succesvol")
	}

	return true
}
