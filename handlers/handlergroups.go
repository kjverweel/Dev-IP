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
	AllGroups := e.Param("groupname")
	GroepID := repositories.GetSepNames
	RecentPosts, err := repositories.GetRecentPosts(GroepID)
	Data := map[string]interface{}{
		"GroupName":   AllGroups,
		"RecentPosts": RecentPosts,
	}
	log.Println(Data)
	return e.Render(http.StatusOK, "sepgroup.html", Data)
}
