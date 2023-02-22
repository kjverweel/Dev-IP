package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
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
		fmt.Println("group already exist or You did fucky")
		return e.Render(http.StatusOK, "groups", nil)

	}
	err = repositories.NewGroup(newGroup)
	if err != nil {
		fmt.Println("Repository got fucked")
	} else {
		fmt.Println("Succesfully called")
	}
	return e.Redirect(http.StatusSeeOther, "/home")
}
