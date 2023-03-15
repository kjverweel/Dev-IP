package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/repositories"
	"net/http"
)

func Groups(e echo.Context) error {
	err := e.Render(http.StatusOK, "groups", nil) //renders de /Groups pagina waar je nieuwe groepen aan kan maken
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()) //als er een error voorkomt, laat hij een http.internalServerError zien met de error message, waarom hij een internal server error geeft
	}
	log.Println("handlergroups.go:Succesfully made it to /groups") //logt een bericht als je succesvol naar een andere pagina bent gegaan
	return nil
}

// dit is de handler voor de aparte groepen die je ziet als je op de links drukt op de homepage/feed
func SepGroup(e echo.Context) error {
	Groepname := e.Param("groupname")                               //e.Param neemt de groupname uit main.go over, en kan daardoor meteen gebruikt worden in de code
	GroepID, err := repositories.GetSepNames(Groepname)             //Groepname string wordt hier via GetSepNames omgezet in een int GroepID
	RecentPosts, err := repositories.GetRecentPosts([]int{GroepID}) //GroepID wordt hier eerst in een slice gezet, omdat dit anders niet werkt met de functie, en daarnadoorgegeven aan GetRecentPosts, dit is een slice van een slice van strings
	if err != nil {
		return err
	}
	return e.Render(http.StatusOK, "sepgroup.html", echo.Map{"GroupName": Groepname, "RecentPosts": RecentPosts}) //als alles goed is gegaan in deze functie hier, rendert dit de: Naam van de groep die bovenaan de pagina staat en de meest recente posts die in deze groep zijn aangemaakt.
}
