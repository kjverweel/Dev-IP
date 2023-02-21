package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Start(e echo.Context) error {
	err := e.Render(http.StatusOK, "index", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
