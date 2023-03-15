package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
)

func Home(e echo.Context) error {
	cookie, err := e.Cookie("User") //get User_ID from cookie
	// parse cookie string value to uint
	userId, err := strconv.ParseUint(cookie.Value, 10, 64)
	if err != nil {
		log.Println("handlerhome.go:Couldn't get cookie")
		e.Render(http.StatusOK, "index", nil) //logt de error en stuurt de user terug aar de login
	}

	user := &models.Users{}                         //Maak een nieuw Users struct aan en wijs het toe aan de user variabele als een pointer
	err = repositories.GetUser(uint(userId), &user) //Roep de GetUser functie aan om gebruikersinformatie op te halen en passt deze toe aan het user struct
	if err != nil {
		log.Println("handlerhome.go: Kon de gebruikersnaam niet ophalen") // Als er een fout optreedt bij het ophalen van de gebruikersinformatie, log dan een foutmelding
	}

	GroepID, err := repositories.GetGroupsFromMembers(int(userId)) //passt userId naar GetGroupsFromMembers, die een slice van ints teruggeeft naar GroepID
	AllGroups := repositories.GetGroups(GroepID)                   //pakt de int's uit de slice van GroepID en zet deze door naar GetGroups, die vervolgens een slice van strings waarin de groepen staan waar de user lid van is
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{ //stuurt een .JSON message met internal server error als er een error voorkomt
			"message": "Failed to get groups",
		})
	}

	if AllGroups == nil { //als er geen groepen zijn, rendert hij de pagina alsnog, alleen met een bericht dat er nog geen groepen zijn. (en omdat je geen posts aan kan maken als er geen groups zijn, hoeven de andere messages er niet in :D)
		e.Render(http.StatusOK, "home", echo.Map{"Groups": "Unfortunately, there are no groups yet"})
	}

	RecentPosts, err := repositories.GetRecentPosts(GroepID) //pakt de GroepID slice van eerder en geeft deze aan de GetRecentPosts repo. Dit haalt de 8 meest recente posts op
	if err != nil {
		log.Println("handlerhome.go:Couldn't get recents posts")
		return nil
	}
	log.Println("handlerhome:", RecentPosts)
	//als alles hier oke is dan rendert hij de homepagina, met daarop de naam van de user die is ingelogd, alle groepen waar hij/zij in zit en de 8 meest recente posts uit alle groepen waar hij/zij in zit
	err = e.Render(http.StatusOK, "home", echo.Map{"Nem": user.UserNickname, "AllGroups": AllGroups, "RecentPosts": RecentPosts})
	if err != nil { //als er een error is dan krijg je een internal server error met daarbij de error
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
