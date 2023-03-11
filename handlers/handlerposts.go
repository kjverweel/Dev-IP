package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/repositories"
	"net/http"
)

func Posts(e echo.Context) error {
	groups, err := repositories.GetGroup()
	log.Println(groups)
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
