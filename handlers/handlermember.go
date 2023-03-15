package handlers

//Ik had deze file ipv handlermember, memberhandler moeten noemen. leuk woordgrapje
import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/repositories"
	"net/http"
)

// handlermember laadt de memberpagina, waar je nieuwe user toe kan voegen aan groepen
func Member(e echo.Context) error {
	AllUsers, err := repositories.GetAllUsers() //hier worden alle users opgevraagd en opgeslagen in AllUsers
	if err != nil {
		return e.Render(http.StatusOK, "home", echo.Map{"Groups": "Unfortunately, We couldn't get the users, so you have to guess the names"})
	} //als hier een error in komt wordt de pagina opnieuw geladen met het een bericht

	groups, err := repositories.GetGroup()
	if err != nil {
		return e.Render(http.StatusOK, "home", echo.Map{"Groups": "Unfortunately, we couldn't get the groups, so you have to guess"})
	} //als hier een error in komt wordt de pagina opnieuw geladen met het een bericht

	if groups == nil {
		e.Render(http.StatusOK, "home", echo.Map{"Groups": "Unfortunately, there are no groups yet"})
	} //als hier een error (ofwel er zijn nog geen groepen) in komt wordt de pagina opnieuw geladen met het een bericht
	err = e.Render(http.StatusOK, "member", echo.Map{"Users": AllUsers, "Groups": groups})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	} //waar de anderen je nog erug sturen krijg je hier een internalserver error. als hier iets fout gaat kan dat niet 123 opgelost worden.
	log.Println("handlermember.go:Succesfully made it to /member")
	return nil
}
