package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
)

func CreateGroup(e echo.Context) error {
	//Cookiecode
	// get cookie from request
	cookie, err := e.Cookie("User")
	// parse cookie string value to uint
	userId, err := strconv.ParseUint(cookie.Value, 10, 64)
	if err != nil {
		//if an error occurs in Cookiecode this usually means that the user isn't logged in properly.
		//this e.Render causes a direct to the index page, where you can log in or register an account.
		log.Println("handlerhome.go:Couldn't get cookie")
		e.Render(http.StatusOK, "index", nil)
	}
	user := &models.Users{}
	err = repositories.GetUser(uint(userId), &user)
	AdminID := strconv.FormatUint(userId, 10)
	log.Println(AdminID)
	if err != nil {
		log.Println("handlerhome.go:Couldn't get cookie")
	}

	newGroup := &models.Groups{
		Groepname:    e.FormValue("Groepsnaam"),
		GroepadminID: AdminID,
	}
	GroupExists, err := repositories.CheckGroup(newGroup)
	if err != nil || GroupExists {
		log.Println("handlercreategroup.go:group already exist or You did fucky")
		return e.Render(http.StatusOK, "groups", echo.Map{"ErrorGroep": "Sorry, deze naam is al in gebruik."})
	}
	err = repositories.NewGroup(newGroup)
	if err != nil {
		log.Println("handlercreategroup.go:Repository got fucked")
	} else {
		log.Println("handlercreategroup.go:Succesfully called")
	}
	return e.Redirect(http.StatusSeeOther, "/home")
}
