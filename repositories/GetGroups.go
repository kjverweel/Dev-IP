package repositories

import "main.go/models"

func GetGroup(groepname **models.Groups) error {
	return db.Where("id = ?", groepname).Error
}
