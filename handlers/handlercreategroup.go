package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
)

func CreateGroup(e echo.Context) error {
	newGroup := &models.Groups{
		Groepname:  e.FormValue("Groepsnaam"),
		Groepadmin: e.FormValue("Adminnaam"),
	}
	GroupExists, err := repositories.CheckGroup(newGroup)
	if err != nil || GroupExists {
		log.Println("group already exist or You did fucky")
		return e.Render(http.StatusOK, "groups", echo.Map{"ErrorGroep": "Sorry, deze naam is al in gebruik."})
	}
	err = repositories.NewGroup(newGroup)
	if err != nil {
		log.Println("Repository got fucked")
	} else {
		log.Println("Succesfully called")
	}
	return e.Redirect(http.StatusSeeOther, "/groepmade")
}
