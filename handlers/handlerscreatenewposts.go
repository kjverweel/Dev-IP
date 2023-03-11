package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
)

func CreateNewPost(e echo.Context) error {

	Groepname := &models.Groups{
		Groepname: e.FormValue("GroupName"),
	}

	GroupID, err := repositories.CompareGroupname(Groepname)
	if err != nil {
		log.Println("handlernewmember.go:couldn't find matching ID")
		return err
	}
	log.Println("handlernewmember.go:GroupID is", GroupID)
	return nil

}