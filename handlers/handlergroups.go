package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/repositories"
	"net/http"
)

func Groups(e echo.Context) error {
	err := e.Render(http.StatusOK, "groups", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Println("handlergroups.go:Succesfully made it to /groups")
	return nil
}

func SepGroup(e echo.Context) error {
	Groepname := e.Param("groupname")
	GroepID, err := repositories.GetSepNames(Groepname)
	log.Println("dit print groepid:", GroepID)
	RecentPosts, err := repositories.GetRecentPosts(GroepID)
	if err != nil {
		return err
	}
	log.Println("Dit print recentposts:", RecentPosts)
	return e.Render(http.StatusOK, "sepgroup.html", echo.Map{"Groupname": Groepname, "RecentPosts": RecentPosts})
}
