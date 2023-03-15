package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/repositories"
	"net/http"
	"strconv"
)

func Posts(e echo.Context) error {
	cookie, err := e.Cookie("User") //get User_ID from cookie
	if err != nil {
		log.Println("couldn't get cookie")
	}
	userId, err := strconv.ParseUint(cookie.Value, 10, 64)
	GroepID, err := repositories.GetGroupsFromMembers(int(userId))
	groups, err := repositories.GetGroup(GroepID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get groups",
		})
	}
	if groups == nil {
		e.Render(http.StatusOK, "home", echo.Map{"Groups": "Unfortunately, there are no groups yet"})
	}
	err = e.Render(http.StatusOK, "newpost", echo.Map{"Groups": groups})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Println("handlerposts.go:Succesfully made it to /newpost")
	return nil
}
