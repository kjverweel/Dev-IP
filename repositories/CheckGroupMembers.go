package repositories

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"main.go/models"
)

func CheckGroupMembers(Groupmembers *models.Groupmembers) (bool, error) {
	var member models.Groupmembers
	err := db.Where("user_id = ?", Groupmembers.UserID).Where("groep_id = ?", Groupmembers.GroepID).First(&member).Error
	if err != nil {
		log.Println(member)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// no matching record found, the user is not yet a member of the group
			return true, nil
		} else {
			// an error occurred while querying the database
			log.Println("error checking group members:", err)
			return false, err
		}
	}
	// a matching record was found, the user is already a member of the group
	return false, errors.New("user is already a member of the group")
}
