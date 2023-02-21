package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Groepjes(e echo.Context) error {
	return e.Render(http.StatusOK, "/groepjes", nil)
}
