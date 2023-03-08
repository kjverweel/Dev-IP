package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
)

func GetNewMemberInfo(e echo.Context) error {

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
	err = e.Render(http.StatusOK, "home", echo.Map{"Groups": groups})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
