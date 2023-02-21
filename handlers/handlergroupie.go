package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Groepjes(e echo.Context) error {
	err := e.Render(http.StatusOK, "groups", nil)
	if err != nil {
		log.Println("Succesfully made it to /groups")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
