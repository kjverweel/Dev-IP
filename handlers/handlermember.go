package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/repositories"
	"net/http"
)

func Member(e echo.Context) error {
	AllUsers, err := repositories.GetAllUsers()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get users",
		})
	}
	groups, err := repositories.GetGroup()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get groups",
		})
	}
	if groups == nil {
		e.Render(http.StatusOK, "home", echo.Map{"Groups": "it no worky worky"})
	}
	err = e.Render(http.StatusOK, "member", echo.Map{"Users": AllUsers})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Println("Succesfully made it to /member")
	return nil
}
