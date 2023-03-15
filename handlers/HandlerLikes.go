package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func GetAndUpdateLikes(e echo.Context) error {
	cookie, err := e.Cookie("User")
	if err != nil {
		log.Println("couldn't get cookie")
		return nil
	}
	UserId, err := strconv.ParseUint(cookie.Value, 10, 64)
	log.Println(UserId)

	return e.Render(http.StatusOK, "home", echo.Map{"LikeCount": UserId})
}

//Needed: Post_ID, User_ID, amount of like already there, once done return like +1