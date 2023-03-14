package repositories

import (
	"log"
	"main.go/models"
)

func IsAnAdmin(UserID int, GroupID int) bool {
	result := &models.Groupmembers{}

	db.Model(&models.Groupmembers{}).Where("user_id = ? AND groep_id = ?", UserID, GroupID).Find(&result)
	log.Println(result.Admin)
	return result.Admin
}
