package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetNewMemberInfo(e echo.Context) error {
	if e.FormValue("UserName") == "" || e.FormValue("GroupName") == "" {
		return e.Render(http.StatusUnauthorized, "member", nil)
	}

	UserNickname := e.FormValue("UserName")
	Groepname := e.FormValue("GroupName")
	log.Println("handlernewmember.go:", UserNickname)
	log.Println("handlernewmember.go:", Groepname)

	return nil
}
