package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
	"time"
)

func Login(e echo.Context) error {
	if e.FormValue("Username") == "" || e.FormValue("Password") == "" {
		return e.Render(http.StatusUnauthorized, "login", nil)
	} //dit stuk hierboven heeft geen nut meer, dit checkt of de Ussername en password uit de form van html zijn. het runnen van dit beetje kost een aantal milliseconden dus ik heb het erin laten staan. in een volledige build haal ik dit er natuurlijk uit

	existingUser := &models.Users{ //slaat de Usernickname en UserPassword uit de forms op in de users struct, in de variable existingUser
		UserNickname: e.FormValue("Username"),
		UserPassword: e.FormValue("Password"),
	}

	YouExist := repositories.LoginUser(existingUser) //hier wordt de existingUser gepasst naar LoginUser. Deze checkt of de ingevoerde username en wachtwoord in de database staan en of ze overeenkomen met wat in de database staat
	if !YouExist {                                   //als de user niet bestaat wordt je terugestuurd naar login met een bericht erbij
		log.Println("handlerlogincheck.go:User doesn't exist")
		return e.Render(http.StatusOK, "login", echo.Map{"UserDoesntExist": "Deze user bestaat niet, probeer opnieuw"})
	}

	e.SetCookie(&http.Cookie{ //als je bent ingelogd krijg je een cookie toegewezen. Deze draagt de naam User, gebruikt de ID uit existingUser en duurt 999 uur voordat deze vergaat. (dit moet in de build ook aangepast worden, 41 dagen is miscchien wat veel :)
		Expires: time.Now().Add(time.Hour * 999),
		Name:    "User",
		Value:   strconv.Itoa(int(existingUser.ID)),
	})
	return e.Redirect(http.StatusSeeOther, "/home") //als laatste als alles goed is wordt de user geredirect naar home
}
