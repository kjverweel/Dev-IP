package handlers

import (
	"github.com/labstack/echo/v4"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
)

func Home(e echo.Context) error {
	// get cookie from request
	cookie, err := e.Cookie("User")
	// parse cookie string value to uint
	userId, err := strconv.ParseUint(cookie.Value, 10, 64)
	if err != nil {
		panic(err)
	}
	user := &models.Users{}
	err = repositories.GetUser(uint(userId), &user)
	if err != nil {
		panic(err)
	}
	err = e.Render(http.StatusOK, "home", echo.Map{"Nem": user.UserNickname})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
