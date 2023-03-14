package repositories

import (
	"github.com/jinzhu/gorm"
	"main.go/models"
)

func IsAnAdmin(CheckForAdmin *models.Groups) (int, error) {
	var IsAdmin string
	err := db.Model(&models.Groups{}).Where("groepname = ? AND groepadmin_id = ?", CheckForAdmin.Groepname, CheckForAdmin.GroepadminID).Pluck("groepadmin_id", &IsAdmin).Error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}
	if IsAdmin == "" {
		return 0, nil
	} else if IsAdmin == CheckForAdmin.GroepadminID {
		return 1, nil
	}
	return 0, nil
}
