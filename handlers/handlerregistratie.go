package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"main.go/models"
	"main.go/repositories"
	"net/http"
)

func Register(e echo.Context) error {
	newUser := &models.Users{
		UserNickname: e.FormValue("Username"),
		UserEmail:    e.FormValue("Email"),
		UserPassword: e.FormValue("Password"),
	}
	YouHere, err := repositories.CompareUsers(newUser)
	if err != nil || YouHere {
		fmt.Println("User already exist or You did fucky")
		return e.Render(http.StatusOK, "register", echo.Map{"NuhUh": "Deze user bestaat al, wees origineel"})

	}
	err = repositories.NewUsers(newUser)
	if err != nil {
		fmt.Println("Repository got fucked")
	} else {
		fmt.Println("Succesfully called")
	}
	return e.Redirect(http.StatusSeeOther, "/home")
}
