package repositories

import (
	"log"
	"main.go/models"
)

func IsAnAdmin(CheckForAdmin models.Groups) (bool, error) {
	var IsAdmin bool
	err := db.Where("groepname = ? AND groepadmin_id = ?", CheckForAdmin.Groepname, CheckForAdmin.GroepadminID).First(&IsAdmin)
	log.Println(IsAdmin)
	if err != nil {
		log.Println("Checkadmin.go: Something doesn't work")
		return false, nil
	}
	if IsAdmin == true {
		return true, nil
	} else if IsAdmin == false {
		return false, nil
	}
	return false, nil
}
