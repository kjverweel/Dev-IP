package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Groups(e echo.Context) error {
	err := e.Render(http.StatusOK, "groups", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Println("Succesfully made it to /groups")
	return nil
}
