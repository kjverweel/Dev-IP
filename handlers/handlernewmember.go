package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/repositories"
	"net/http"
)

func Member(e echo.Context) error {
	groups, err := repositories.GetGroup()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get groups",
		})
	}
	err = e.Render(http.StatusOK, "member", echo.Map{"Groups": groups})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Println("Succesfully made it to /groups")
	return nil
}
