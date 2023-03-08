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

	return nil
}
