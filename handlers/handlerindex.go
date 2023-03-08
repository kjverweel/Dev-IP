package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Loginscreen(e echo.Context) error {
	err := e.Render(http.StatusOK, "login", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		log.Println("handlerindex.go:Succesfully made it to login!")
	}
	return nil
}

func Registerscreen(e echo.Context) error {
	err := e.Render(http.StatusOK, "register", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		log.Println("handlerlogin.go:Succesfully made it to register!")
	}
	return nil
}
