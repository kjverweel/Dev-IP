package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
	"time"
)

func Login(e echo.Context) error {
	if e.FormValue("Username") == "" || e.FormValue("Password") == "" {
		return e.Render(http.StatusUnauthorized, "login", nil)
	} //the above bit is no longer necessary, but since it's inactive and costs a few milliseconds to run i'm keeping it for reference

	existingUser := &models.Users{
		UserNickname: e.FormValue("Username"),
		UserPassword: e.FormValue("Password"),
	}
	log.Println(existingUser)

	YouExist := repositories.LoginUser(existingUser)
	if !YouExist {
		log.Println("User doesn't exist")
		return e.Render(http.StatusOK, "login", echo.Map{"UserDoesntExist": "Deze user bestaat niet, probeer opnieuw"})
	}

	e.SetCookie(&http.Cookie{
		Expires: time.Now().Add(time.Hour * 999),
		Name:    "User",
		Value:   strconv.Itoa(int(existingUser.ID)),
	})
	return e.Redirect(http.StatusSeeOther, "/home")
}
