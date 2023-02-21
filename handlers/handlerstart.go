package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Start(e echo.Context) error {
	err := e.Render(http.StatusOK, "index", nil)
	if err != nil {
		log.Println("Succesfully made it to /home")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
