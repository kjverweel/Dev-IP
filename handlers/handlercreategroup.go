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
	cookie, err := e.Cookie("User") //get User_ID from cookie
	if err != nil {
		log.Println("couldn't get cookie")
	}
	UserId, err := strconv.ParseUint(cookie.Value, 10, 64) //convert from cookie
	log.Println("Handlercreategroup.go:", UserId)
	if err != nil {
		log.Println("handlerhome.go:Couldn't get cookie")
	}

	newGroup := &models.Groups{
		Groepname: e.FormValue("Groepsnaam"),
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
	NewGroupID := repositories.GetLatestGroup()
	log.Println(NewGroupID)

	NewMember := &models.Groupmembers{
		UserID:  int(UserId),
		GroepID: NewGroupID,
		Admin:   true,
	}
	err = repositories.NewMember(NewMember)
	return e.Redirect(http.StatusSeeOther, "/home")
}
