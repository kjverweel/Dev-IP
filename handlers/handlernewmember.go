package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
)

func GetNewMemberInfo(e echo.Context) error {
	//Cookiecode
	// get cookie from request
	cookie, err := e.Cookie("User")
	// parse cookie string value to uint
	userId, err := strconv.ParseUint(cookie.Value, 10, 64)
	if err != nil {
		//if an error occurs in Cookiecode this usually means that the user isn't logged in properly.
		//this e.Render causes a direct to the index page, where you can log in or register an account.
		log.Println("handlerhome.go:Couldn't get cookie")
		e.Render(http.StatusOK, "index", nil)
	}
	user := &models.Users{}
	err = repositories.GetUser(uint(userId), &user)
	if err != nil {
		log.Println("handlerhome.go:Couldn't get cookie")
	}
	//end of Cookiecode

	if e.FormValue("UserName") == "" || e.FormValue("GroupName") == "" {
		return e.Render(http.StatusUnauthorized, "member", nil)
	}

	Usernickname := &models.Users{
		UserNickname: e.FormValue("UserName"),
	}
	Groepname := &models.Groups{
		Groepname: e.FormValue("GroupName"),
	}

	log.Println("handlernewmember.go:", Usernickname)
	log.Println("handlernewmember.go:", Groepname)

	UserID, err := repositories.CompareUsername(Usernickname)
	if err != nil {
		log.Println("handlernewmember.go:couldn't find matching ID")
		return err
	}

	log.Println("handlernewmember.go:UserID is", UserID)

	groups, err := repositories.GetGroup()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get groups",
		})
	}
	if groups == nil {
		e.Render(http.StatusOK, "home", echo.Map{"Groups": "Unfortunately, there are no groups yet"})
	}
	err = e.Render(http.StatusOK, "home", echo.Map{"Nem": user.UserNickname, "Groups": groups})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
