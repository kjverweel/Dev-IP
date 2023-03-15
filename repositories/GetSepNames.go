package repositories

import (
	"log"
	"main.go/models"
)

func GetSepNames(Groepname string) (int, error) {
	group := &models.Groups{}
	db.Where("groepname = ?", Groepname).First(&group)
	log.Println(group)
	return int(group.ID), nil
}
