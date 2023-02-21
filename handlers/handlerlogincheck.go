package handlers

import (
	"fmt"
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
	}
	existingUser := &models.Users{
		UserNickname: e.FormValue("Username"),
		UserPassword: e.FormValue("Password"),
	}

	log.Println(existingUser)
	
	YouExist := repositories.LoginUser(existingUser)
	if !YouExist {
		fmt.Println("User doesn't exist")
		return e.Render(http.StatusOK, "login", echo.Map{"UserDoesntExist": "Deze user bestaat niet, probeer opnieuw"})
	}
	e.SetCookie(&http.Cookie{
		Expires: time.Now().Add(time.Hour * 999),
		Name:    "User",
		Value:   strconv.Itoa(int(existingUser.ID)),
	})
	return e.Redirect(http.StatusSeeOther, "/home")
}
