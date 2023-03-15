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
		return e.Render(http.StatusUnauthorized, "groups", echo.Map{"ErrorGroep": "sorry, maar momenteel kan je geen groepen maken. Contact een admin of kom later terug."}) //logs the error and return user to login
	}
	UserId, err := strconv.ParseUint(cookie.Value, 10, 64) //convert cookie from struct field to Unique integer
	if err != nil {
		log.Println("handlerhome.go:Couldn't get cookie")
		return e.Render(http.StatusNotFound, "groups", echo.Map{"ErrorGroep": "sorry, maar momenteel kan je geen groepen maken. Contact een admin of kom later terug."}) // logs the error and returns to the /groups page
	}
	newGroup := &models.Groups{
		Groepname: e.FormValue("Groepsnaam"),
	} // newGroup slaat de gegevens uit het HTML formulier op in de models.Groups struct

	GroupExists, err := repositories.CheckGroup(newGroup) //newGroup wordt hier gebruikt om de data uit het formulier door te geven aan repositories.CheckGroup
	if err != nil || GroupExists {                        //als er een error is, of de groep bestaat, logt het een error, en stuurt hij je terug naar /groups
		log.Println("handlercreategroup.go:group already exist or You did fucky")
		return e.Render(http.StatusOK, "groups", echo.Map{"ErrorGroep": "Sorry, deze naam is al in gebruik."})
	}
	err = repositories.NewGroup(newGroup) //als de groep niet bestaat word deze hier aangemaakt. door gebruik van gorm hoef je hier alleen de naam in te vullen
	if err != nil {
		log.Println("handlercreategroup.go:Repository couldn't get called")
		return e.Render(http.StatusNotFound, "groups", echo.Map{"ErrorGroep": "sorry, maar momenteel kan je geen groepen maken. Contact een admin of kom later terug."})
	} else {
		log.Println("handlercreategroup.go:Succesfully called")
	}
	NewGroupID := repositories.GetLatestGroup() //hier pakt GetLatesttGroups, direct nadat de groep is aangemaakt, het ID van deze groep en zet het in de variable NewGroupID

	NewMember := &models.Groupmembers{ //NewMember wijst UserID, GroepID en Admin = true aan naar de Groupmembers struct. Admin staat hier op True omdat deze struct alleen wordt aangeroepen wanneer er een groep wordt aangemaakt. dit maakt degene die de groep aanmaakt automatisch lid en admin van de groep
		UserID:  int(UserId),
		GroepID: NewGroupID,
		Admin:   true,
	}
	err = repositories.NewMember(NewMember)         //stuurt de NewMember gegevens door naar de NewMember repository waar de nieuwe admin wordt toegevoegd
	return e.Redirect(http.StatusSeeOther, "/home") //zodra alles klaar is, redirect hij je terug naar home, en zie je de nieuwe groep staan
}
